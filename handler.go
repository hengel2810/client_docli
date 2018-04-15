package main

import (
	"github.com/hengel2810/client_docli/config"
	"fmt"
	"github.com/hengel2810/client_docli/login"
	"flag"
	"os"
	"github.com/hengel2810/client_docli/docker"
	"github.com/hengel2810/client_docli/api"
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
		strImageId := *imageId
		uploadImage, err := docker.UploadDockerImage(strImageId)
		if err != nil {
			fmt.Println(err)
		} else {
			api.PostImageData(uploadImage)
		}
	}
}