package downloader

import (
	"github.com/shivshankarsingh98/file-downloader/internal/database"
	"log"
)

type Downloader interface {
	DownloadFile()
	StoreLocal()
}

func NewDownloader(protocol string, db database.Database) Downloader {
	config := db.GetConfig()
	// this function return the struct for the protocol needed for this instance
	switch protocol {
	case "ftp":
		return newFtp(config)
	case "http":
		return newHttp(config)
	case "sftp":
		return newSftp(config)
	default:
		log.Println("protocol not defined: ",protocol)

	}
	return nil
}