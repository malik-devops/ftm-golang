package database

import "fmt"

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	Port       string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.Port, config.DB)

	return connectionString
}
