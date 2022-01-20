package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	NoSSl = "disable"
)

type PostgresClient struct {
	db       *sql.DB
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	SslMode  string
}

func (pg *PostgresClient) IitDbConn() error {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=%s",
		pg.User, pg.Password, pg.Host, pg.Port, pg.DbName, pg.SslMode)
	db, err := sql.Open("postgres", connStr)
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

//func (pg *PostgresClient) GetAdList() ([]Ad, error) {
//	if pg.db == nil {
//		return nil, errors.New(noConnectionErr)
//	}
//	rows, err := pg.db.Query("select a.id, a.name, a.description, a.price , a.create_ts from public.ad a")
//	if err != nil {
//		return nil,err
//	}
//	adResult := make([]Ad,0,0)
//
//	for rows.Next() {
//		var ad Ad
//		err := rows.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price, &ad.CreateTs)
//		if err != nil {
//			return nil,err
//		}
//		adResult = append(adResult, ad)
//	}
//	return adResult, nil
//}
