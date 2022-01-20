package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	NoSSl = "disable"
)

type PostgresClient struct {
	db     *sql.DB
	DbConn string
}

func (pg *PostgresClient) IitDbConn() error {
	db, err := sql.Open("postgres", pg.DbConn)
	if err != nil {
		return err
	}

	pg.db = db
	return nil
}

func (pg *PostgresClient) CloseDbConn() error {
	if pg.db != nil {
		return pg.db.Close()
	}
	return nil
}

func (pg *PostgresClient) ExecuteQuery(query string,
	rowsParser func(row *sql.Rows) error) error {
	if pg.db == nil {
		if pg.db == nil {
			return NoConnectionErr
		}
	}
	if query == "" {
		return EmptyQuery
	}
	rows, err := pg.db.Query(query)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err = rowsParser(rows); err != nil {
			return err
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Close(); err != nil {
		return err
	}
	return nil
}
