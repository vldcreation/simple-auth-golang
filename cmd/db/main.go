package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
)

func main() {
	env := entity.ENV()

	var (
		pgUser     = env.Get(constants.DBUser)
		pgPassword = env.Get(constants.DBPwd)
		pgHost     = env.Get(constants.DBHost)
		pgPort     = env.Get(constants.DBPort)
		pgDbName   = env.Get(constants.DBName)
		ctx        = context.TODO()
		xsql       = new(repository.SQL)
		conn       = repository.SQLConn(nil)
		dsn        = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPassword, pgHost, pgPort, pgDbName)
	)

	fmt.Println("dsn:", dsn)

	conn = xsql.WithDSN(dsn)

	queries := []string{
		`drop schema if exists edufund;`,
		`create schema IF NOT EXISTS edufund;`,
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
		`CREATE UNIQUE INDEX ON edufund.users(username);`,
		`CREATE UNIQUE INDEX ON edufund.users(email);`,
	}

	for _, query := range queries {
		err := xsql.SetupOrTeardown(ctx, conn,
			query)
		if err != nil {
			log.Printf("Error: %v", err)
			panic(err)
		}
	}

	log.Default().Println("SetupDB Done")
}
