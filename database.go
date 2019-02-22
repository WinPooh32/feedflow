package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func initDatabse(setts settings, debug bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch driver := *setts.DbDriver; driver {
	case "postgres":
		db, err = initPostgres(setts)
	default:
		return nil, fmt.Errorf("Unsupported database:%s", driver)
	}

	if err != nil {
		return nil, err
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
