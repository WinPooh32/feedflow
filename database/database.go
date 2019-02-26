package database

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	postgres = "postgres"
	mysql    = "mysql"
	sqlite   = "sqlite3"
)

//Credential - database connection information
type Credential struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Database string
	Password string
	Ssl      bool
}

func initPostgres(cred Credential) (*gorm.DB, error) {
	var sslmode string
	if cred.Ssl {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	dbArgs := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cred.Host, cred.Port, cred.User,
		cred.Database, cred.Password, sslmode)

	db, err := gorm.Open(cred.Driver, dbArgs)
	return db, err
}

func initSqlite(cred Credential) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite, cred.Database+"."+sqlite)
	return db, err
}

//Init - setup database connection if possible
func Init(cred Credential, debug bool) (db *gorm.DB, err error) {
	//Try to use any db if we can
	switch driver := cred.Driver; driver {
	case postgres:
		db, err = initPostgres(cred)
	// case "mysql":
	//TODO mysql initialization
	case sqlite:
		db, err = initSqlite(cred)
	default:
		return nil, fmt.Errorf("Unsupported or unavailable database:%s", driver)
	}

	if err != nil {
		//try to fallback to sqlite database
		if cred.Driver != sqlite {
			db, err = initSqlite(cred)
		}

		if err != nil {
			return nil, err
		}
	}

	log.Println("Successfully connected to DataBase!")

	if debug {
		db.LogMode(true)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	//Set ORM table naming
	db.SingularTable(true)

	return db, nil
}

//NewMiddleware creates new gin middleware for sharing database object
func NewMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
	}
}

//Extract - try to extract database object from gin context
func Extract(ctx *gin.Context) (db *gorm.DB, ok bool) {
	dbInterface, _ := ctx.Get("database")
	db, ok = dbInterface.(*gorm.DB)
	return db, ok
}
