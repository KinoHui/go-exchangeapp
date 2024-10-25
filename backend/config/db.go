package config

import (
	"exchangeapp/backend/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Fail to initializedatabase, got error: %v", err)
	}

	dbSql, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}

	dbSql.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	dbSql.SetMaxOpenConns(AppConfig.Database.MaxOpenCons)
	dbSql.SetConnMaxLifetime(time.Hour)

	global.Db = db

}
