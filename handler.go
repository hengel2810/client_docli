package main

import (
	"github.com/hengel2810/client_docli/config"
	"fmt"
	"github.com/hengel2810/client_docli/login"
	"github.com/hengel2810/client_docli/fs"
	"github.com/hengel2810/client_docli/controller"
	"github.com/hengel2810/client_docli/docker"
	"flag"
	"os"
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

func HandleUploadFromConfig() {
	uploadCommand := flag.NewFlagSet("start", flag.ExitOnError)
	configFile := uploadCommand.String("config", "", "config file (.yml or .yaml)")
	uploadCommand.Parse(os.Args[2:])
	if uploadCommand.Parsed() {
		if *configFile == "" {
			fmt.Println("Please supply config file")
			return
		}
		strConfigFile := *configFile
		docli, err := fs.ReadConfig(strConfigFile)
		if err != nil {
			fmt.Println(err)
		} else {
			docli, err = controller.SetDocliObjectData(docli)
			if err != nil {
				fmt.Println(err)
			} else {
				err = docker.UploadDockerImage(docli)
				if err != nil {
					fmt.Println(err)
				} else {
					api.PostImageData(docli)
					fmt.Println("Image sucessfully pushed")
				}
			}
		}
	}
}