package database

import (
	"database/sql"
)

var postgresClient *sql.DB

type postgress struct {
	client *sql.DB
}

func newPostgres() *postgress {
	// this function return new postgres instance
	return &postgress{
		client: getDbconn(),
	}
}

func (p *postgress) GetConfig() Config {
	//this method gets the Config details of the file to download from postgres database
	return Config{}
}

func getDbconn() *sql.DB {
	//this function opens only a single instance of postgress client connection
	if postgresClient == nil {
		if postgresClient == nil {
			postgresClient, _ = sql.Open("driverName", "dataSourceName")
			return postgresClient
		}
		return postgresClient
	}
	return postgresClient
}
