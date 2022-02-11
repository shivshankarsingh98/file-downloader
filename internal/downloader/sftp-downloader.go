package downloader

import (
	"github.com/shivshankarsingh98/file-downloader/internal/database"
	"log"
)

type sftp struct {
	config database.Config
}

func newSftp(config database.Config) *sftp {
	//this method create new struct instance for sftp protocol
	return &sftp{
		config: config,
	}
}

func (s *sftp) DownloadFile() {
	// this method download the file using its config file

	log.Println("downloading the file using sftp protocol")
}

func (s *sftp) StoreLocal() {
	//this method, after downloading stores the file locally

	log.Println("file stored in local using sftp protocol")
}