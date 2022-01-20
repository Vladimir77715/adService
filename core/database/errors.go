package database

import "errors"

var (
	NoConnectionErr = errors.New("connection is not establish")
	EmptyQuery      = errors.New("empty query request")
)
