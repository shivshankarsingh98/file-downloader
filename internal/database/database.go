package database

import "log"

type Database interface {
	GetConfig() Config

}

type Config struct {
	Protocol   string
	RemoteHost string
	Username   string
	Password   string
}

func NewDatabase(dbName string) Database {
	switch dbName {
	case "postgres":
		return newPostgres()
	default:
		log.Println("unknown db: ",dbName)
	}
	return nil
}