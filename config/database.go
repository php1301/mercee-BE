package config

import "fmt"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBname   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     Env("DB_HOST"),
		Port:     Env("DB_PORT"),
		User:     Env("DB_USER"),
		DBname:   Env("DB_NAME"),
		Password: Env("DB_PASSWORD"),
	}
	return &dbConfig
}

func DbUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBname,
	)
}
