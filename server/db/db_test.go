package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "plants",
		User:       "admin",
		Password:   "1111",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://admin:1111@localhost/plants?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
