package postgresql_query

import (
	_ "embed"
)

var (
	// file format => {SCHEMA}.{TABLE}--{COMMAND}[.{EXTRA}].sql
	///////////////////////////////////////////////////////////////////////////.

	//go:embed edufund.users--add_new_user.insert.sql
	AddNewUser string
	//go:embed edufund.users--get_user_by_id.select.sql
	GetUserByID string
	//go:embed edufund.users--get_user_by_email_or_username.select.sql
	GetUserByEmailOrUsername string
)
