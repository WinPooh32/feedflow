package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	postgresql = "postgres"
	mysql      = "mysql"
	sqlite     = "sqlite3"
)

func initPostgres(setts settings) (*gorm.DB, error) {
	var sslmode string
	if *setts.DbSsl {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	dbArgs := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		*setts.DbHost, *setts.DbPort, *setts.DbUser,
		*setts.DbName, *setts.DbPassword, sslmode)

	db, err := gorm.Open(*setts.DbDriver, dbArgs)
	return db, err
}

func initSqlite(setts settings) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite, *setts.DbName+"."+sqlite)
	return db, err
}

func initDatabse(setts settings, debug bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	//Try to use any db if we can
	switch driver := *setts.DbDriver; driver {
	case postgresql:
		db, err = initPostgres(setts)
	// case "mysql":
	//TODO mysql initialization
	case sqlite:
		db, err = initSqlite(setts)
	default:
		return nil, fmt.Errorf("Unsupported or unavailable database:%s", driver)
	}

	if err != nil {
		//try to fallback to sqlite database
		if *setts.DbDriver != sqlite {
			db, err = initSqlite(setts)
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

func databaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
	}
}
