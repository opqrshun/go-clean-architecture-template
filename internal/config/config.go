package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type (
	AppConfig struct {
		Env       string `env:"ENV",notEmpty`
		SentryDSN string `env:"SENTRY_DSN",notEmpty`
	}

	DBConfig struct {
		User     string `env:"MYSQL_USER",notEmpty`
		Pass     string `env:"MYSQL_PASSWORD",notEmpty`
		Protocol string `env:"MYSQL_PROTOCOL",notEmpty`
		Database string `env:"MYSQL_DATABASE",notEmpty`
	}

	ExtAPIConfig struct {
		URL string `env:"EXTERNAL_API_URL",notEmpty`
	}
)

var appConfig *AppConfig
var dbConfig *DBConfig
var extAPIConfig *ExtAPIConfig

func init() {
	log.Print("loading .env file")
	err := godotenv.Load()

	if err != nil {
		log.Printf("ERR: failed to load .env file :%v", err)
	}

	appConfig = NewAppConfig()
	dbConfig = NewDBConfig()
	extAPIConfig = NewExtAPIConfig()

}

func isLocal() bool {
	_env, ok := os.LookupEnv("ENV")
	if !ok {
		return true
	}
	return _env == "local"
}

func parseEnv(config interface{}) {
	if err := env.Parse(config); err != nil {
		log.Printf("ERR: failed to parse env :%v", err)
		if !isLocal() {
			panic(err.Error())
		}
	}
}

func NewAppConfig() *AppConfig {
	log.Print("loading AppConfig")
	var c AppConfig
	parseEnv(&c)

	return &c
}

func NewDBConfig() *DBConfig {
	var c DBConfig

	log.Print("loading DBConfig")

	parseEnv(&c)

	return &c
}

func NewExtAPIConfig() *ExtAPIConfig {
	log.Print("loading ExtAPIConfig")
	var c ExtAPIConfig
	parseEnv(&c)

	return &c
}

func GetAppConfig() *AppConfig {
	return appConfig
}

func GetDBConfig() *DBConfig {
	return dbConfig
}

func GetExtAPIConfig() *ExtAPIConfig {
	return extAPIConfig
}
