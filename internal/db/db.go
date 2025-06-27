package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type DSNBuilder interface {
	BuildDSN() string
}

type Database struct {
	DB      *gorm.DB
	Dialect string
}

func Connect() (*Database, error) {

	driver := os.Getenv("DB_DRIVER")
	dsnB, DSNErr := GetDSNBuilder(driver)
	if DSNErr != nil {
		return nil, DSNErr
	}
	dsn := dsnB.BuildDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	maxIdleConns := getEnvAsInt("DB_MAX_IDLE_CONNS", 5)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	maxOpenConns := getEnvAsInt("DB_MAX_OPEN_CONNS", 10)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	connMaxLifetime := getEnvAsInt("DB_CONN_MAX_LIFETIME", 300)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	fmt.Println("GORM database connection established successfully")

	return &Database{DB: db, Dialect: driver}, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get generic database object: %w", err)
	}
	return sqlDB.Close()
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := os.Getenv(name)
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
