package database

import "database/sql"

type DbClient interface {
	IitDbConn() error
	CloseDbConn() error
	ExecuteQuery(query string, rowsParser func(row *sql.Rows) error) error
}
