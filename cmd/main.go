package main

import (
	"github.com/shivshankarsingh98/file-downloader/internal/database"
	"github.com/shivshankarsingh98/file-downloader/internal/downloader"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main(){
	//databse interface for DB operations
	database := database.NewDatabase("postgres")

	// gets the list of protocols this instance is responsible from config.yaml
	// and then downloads and stores the file locally
	for _, protocol := range getProtocols() {
		downloader := downloader.NewDownloader(protocol,database) // Downloader interface for download operations
		downloader.DownloadFile()
		downloader.StoreLocal()
	}

}

func getProtocols() []string {
	// this function reads config.yaml file for the list of protocol this instance should download the file

	yamlFile, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	protocols := map[string][]string{}
	err = yaml.Unmarshal(yamlFile, &protocols)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return protocols["protocols"]
}