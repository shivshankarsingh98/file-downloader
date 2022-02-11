package downloader

import (
	"github.com/shivshankarsingh98/file-downloader/internal/database"
	"log"
)

type ftp struct {
	config database.Config
}

func newFtp(config database.Config) *ftp {
	//this method create new struct instance for ftp protocol
	return &ftp{
		config: config,
	}
}

func (f *ftp) DownloadFile() {
	// this method download the file using its config file

	log.Println("downloading the file using ftp protocol")
}

func (f *ftp) StoreLocal() {
	//this method, after downloading stores the file locally

	log.Println("file stored in local using ftp protocol")
}