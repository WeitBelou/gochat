package config

import "net/url"

type DB struct {
	Host   string
	DBName string

	User     string
	Password string
}

func (d *DB) ToPostgresDSN() string {
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(d.User, d.Password),
		Host:     d.Host,
		Path:     d.DBName,
		RawQuery: "sslmode=disable",
	}
	return u.String()
}