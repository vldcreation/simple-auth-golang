package constants

import "github.com/jmoiron/sqlx"

var (
	DB *sqlx.DB

	AppKey  = "APP_KEY"
	AppName = "APP_NAME"
	AppHost = "APP_HOST"
	AppPort = "APP_PORT"

	NameSpace = "APP_NAMESPACE"
	DBEngine  = "DB_ENGINE"
	DBHost    = "DB_HOST"
	DBPort    = "DB_PORT"
	DBUser    = "DB_USER"
	DBPwd     = "DB_PWD"
	DBName    = "DB_NAME"
	DBSSLMode = "DB_SSL_MODE"
	DBConnStr = "DB_CONN_STR"
)
