package app

import (
	"context"
	"time"

	"github.com/gearintellix/u2"
	"github.com/jmoiron/sqlx"
	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/utils/utstring"

	_ "github.com/lib/pq"
)

func (ox *App) initDB(ctx context.Context) (err error) {
	env := entity.ENV()
	sqlConn := utstring.Env(constants.DBConnStr, `
		host=__host__
		user=__user__
		password=__pwd__
		dbname=__name__
		sslmode=__sslMode__
		application_name=__appKey__
		port=__port__
	`)
	sqlConn = u2.Binding(sqlConn, map[string]string{
		"host":    env.Get(constants.DBHost),
		"user":    env.Get(constants.DBUser),
		"pwd":     env.Get(constants.DBPwd),
		"name":    env.Get(constants.DBName),
		"sslMode": env.Get(constants.DBSSLMode),
		"appKey":  env.Get(constants.AppKey),
		"appName": env.Get(constants.AppName),
		"port":    env.Get(constants.DBPort),
	})

	ox.DB, err = sqlx.Connect("postgres", sqlConn)
	if err != nil {
		return err
	}

	ox.DB.SetConnMaxLifetime(time.Second * 14400)

	constants.DB = ox.DB

	return ox.DB.PingContext(ctx)
}

func (ox *App) getDB(ctx context.Context) *sqlx.DB {
	return ox.DB
}
