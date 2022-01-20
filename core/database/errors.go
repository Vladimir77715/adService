package database

import "errors"

var (
	NoConnectionErr = errors.New("connection is not establish")
	ConnEnvEmpty    = errors.New("environment CONN_DB is empty")
	EmptyQuery      = errors.New("empty query request")
)
