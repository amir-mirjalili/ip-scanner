package db

import (
	"fmt"
	"os"
	"strconv"
)

var DSNRegistry = make(map[string]DSNBuilder)

func RegisterDSNBuilder(dsn string, builder DSNBuilder) {
	DSNRegistry[dsn] = builder
}

func GetDSNBuilder(dsn string) (DSNBuilder, error) {
	builder, exists := DSNRegistry[dsn]
	if !exists {
		return nil, fmt.Errorf("database driver %s is not registered", dsn)
	}
	return builder, nil
}

func init() {
	RegisterDSNBuilder("mysql", &MySQLDSNBuilder{})
}

type MySQLDSNBuilder struct{}

func (m *MySQLDSNBuilder) BuildDSN() string {
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 3306
	}
	password := os.Getenv("DB_PASSWORD")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=%s",
		os.Getenv("DB_USER"),
		password,
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)
}
