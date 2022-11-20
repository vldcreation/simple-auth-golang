package app

import (
	"context"
	"log"
	"time"

	"github.com/gearintellix/u2"
	"github.com/jmoiron/sqlx"
	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
	"github.com/vldcreation/simple-auth-golang/internal/utils/utstring"

	_ "github.com/lib/pq"
)

func (ox *App) SetupOrTeardown(ctx context.Context) {

}

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

	queries := []string{
		`CREATE schema IF NOT EXISTS edufund;`,
		`CREATE TABLE IF NOT EXISTS edufund.users (
			id bigserial primary key,
			fullname varchar(255) not null,
			email varchar(100) not null,
			username varchar(100) not null,
			"password" varchar(255) not null,
			created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
			created_at timestamptz NOT NULL DEFAULT now(),
			modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
			modified_at timestamptz NOT NULL DEFAULT now(),
			deleted_by varchar(255) NULL,
			deleted_at timestamptz NULL
		);`,
		`DO $$
		BEGIN
			IF (SELECT to_regclass('edufund.users') IS NOT null) 
			THEN 
				CREATE INDEX IF NOT EXISTS users_username_idx
				ON edufund.users(username)
				TABLESPACE pg_default;
			END IF;
		END $$;`,
		`DO $$
		BEGIN
			IF (SELECT to_regclass('edufund.users') IS NOT null) 
			THEN 
				CREATE INDEX IF NOT EXISTS users_email_idx
				ON edufund.users(email)
				TABLESPACE pg_default;
			END IF;
		END $$;`,
	}

	var xsql = new(repository.SQL)

	err = xsql.SetupOrTeardown(ctx, ox.DB,
		queries[0])
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	err = xsql.SetupOrTeardown(ctx, ox.DB,
		queries[1])
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	err = xsql.SetupOrTeardown(ctx, ox.DB,
		queries[2])
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	err = xsql.SetupOrTeardown(ctx, ox.DB,
		queries[3])
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	log.Default().Println("SetupDB Done")

	ox.DB.SetConnMaxLifetime(time.Second * 14400)

	constants.DB = ox.DB

	return ox.DB.PingContext(ctx)
}

func (ox *App) getDB(ctx context.Context) *sqlx.DB {
	return ox.DB
}
