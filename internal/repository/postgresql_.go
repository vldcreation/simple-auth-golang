package repository

type postgresql struct{ tc SQLTxConn }

type PostgreSQL interface {
	AddNewUser
	GetUserByID
	GetUserByEmailOrUsername
}

func NewPostgreSQL(txc SQLTxConn) PostgreSQL { return &postgresql{txc} }
