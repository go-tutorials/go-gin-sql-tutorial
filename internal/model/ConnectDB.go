package model

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "sqlserver://sa:admin@123456@localhost:1433?database=demogo"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&Users{})

	database, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	database.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	database.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	database.SetConnMaxLifetime(time.Hour)

	return db
}
