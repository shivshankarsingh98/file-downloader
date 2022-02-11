# file-downloader

App entry point: file-downloader/cmd/main.go

Steps to run the app
    1) clone the repo
    2) Go to project folder file-downloader
            cd file-downloader

    3) Download the dependency
            go mod download
    4) Add the protocol in config.yaml file , for which you want this instance to download the files
            Example yaml file
            
            `
            protocols:
              - ftp
            #  - http
            #  - sftp
            `
            
    5) run the app
            cd cmd 
            go run main.go
    
