package main

import (
	"github.com/hengel2810/client_docli/config"
	"fmt"
	"github.com/hengel2810/client_docli/login"
	"flag"
	"os"
	"github.com/hengel2810/client_docli/api"
	"github.com/hengel2810/client_docli/docker"
)

func HandleLogin()  {
	configValid := config.ConfigValid()
	if !configValid {
		login.StartLoginServer()
	} else {
		fmt.Println("Already logged in. Please use 'docli logout' to logout before re-login")
	}
}

func HandleLogout() {
	config.RemoveTokenConfig()
}

func HandleUpload() {
	uploadCommand := flag.NewFlagSet("upload", flag.ExitOnError)
	imageId := uploadCommand.String("image", "", "image id or name")
	uploadCommand.Parse(os.Args[2:])
	if uploadCommand.Parsed() {
		if *imageId == "" {
			fmt.Println("Please supply image id or name")
			return
		}
		tempPath := os.TempDir()
		fileName := *imageId + ".tar"
		filePath := tempPath + fileName
		copied := docker.CopyImage(*imageId, filePath)
		if copied {
			postError, statusCode := api.PostFile(fileName, filePath, "http://localhost:8000/image", *imageId)
			if postError == nil && statusCode == 200 {
				fmt.Println("Image uploaded")
			} else {
				fmt.Println("Error while uploading image")
			}
		} else {
			fmt.Println("Error while copying image")
		}
	}
}