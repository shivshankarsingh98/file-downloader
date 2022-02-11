package downloader

import (
	"github.com/shivshankarsingh98/file-downloader/internal/database"
	"log"
)

type http struct {
	config database.Config
}

func newHttp(config database.Config) *http {
	//this method create new struct instance for http protocol
	return &http{
		config: config,
	}
}

func (h *http) DownloadFile() {
	// this method download the file using its config file

	log.Println("downloading the file using http protocol")
}

func (h *http) StoreLocal() {
	//this method, after downloading stores the file locally

	log.Println("file stored in local using http protocol")
}