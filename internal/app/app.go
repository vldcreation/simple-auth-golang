package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/service/delivery"
)

type App struct {
	DB        *sqlx.DB
	Hostname  string
	GinObject delivery.GinObject
}

func Run(ctx context.Context) {
	env := entity.ENV()

	require := []string{
		"APP_KEY",
		"APP_NAME",
		"APP_HOST",
		"APP_PORT",
		"DB_ENGINE",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PWD",
		"DB_NAME",
		"DB_SSL_MODE",
		"SERVER_HOST_NAME",
		"SERVER_MAIL_ADDRESS",
		"JWT_SECRET_KEY",
	}

	app := new(App)
	app.Hostname, _ = os.Hostname()

	if err := env.Require(require...); err != nil {
		log.Fatalln("missing environment variables", require)
		<-time.After(time.Second * 5)
		panic(err)
	}

	if err := app.initDB(ctx); err != nil {
		log.Fatalln("app.initDB();\n", err)
		<-time.After(time.Second * 5)
		panic(err)
	}

	if err := app.initService(ctx); err != nil {
		log.Fatalln("app.initService();")
		<-time.After(time.Second * 5)
		panic(err)
	}

	app.GinObject.InitRoutes(ctx)
}
