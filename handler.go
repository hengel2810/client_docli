package main

import (
	"github.com/hengel2810/client_docli/config"
	"fmt"
	"github.com/hengel2810/client_docli/fs"
	"github.com/hengel2810/client_docli/controller"
	"github.com/hengel2810/client_docli/docker"
	"flag"
	"os"
	"github.com/hengel2810/client_docli/api"
	"github.com/crackcomm/go-clitable"
)

func HandleLogin()  {
	configValid := config.ConfigValid()
	if !configValid {
		api.StartLoginProcess()
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
					err := api.PostImageData(docli)
					if err != nil {
						fmt.Println(err)
						fmt.Println("Error pushing image")
					} else {
						fmt.Println("Image sucessfully pushed")
					}
				}
			}
		}
	}
}

func HandleListDoclis() {
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		fmt.Println("Error loading Token")
	} else {
		arrDocli, err := api.GetDoclis(cfg.UserId)
		if err != nil {
			fmt.Println("Error listing doclis")
		} else {
			table := clitable.New([]string{"Name", "Image", "Created", "Id"})
			for _, docli := range arrDocli {
				table.AddRow(map[string]interface{}{
					"Name": docli.ContainerName,
					"Image": docli.OriginalName,
					"Created": docli.Uploaded.Local().String(),
					"Id": docli.UniqueId,
				})
			}
			table.Markdown = true
			table.Print()
		}
	}
}

func HandleRemoveDocli() {
	removeCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	docliId := removeCommand.String("id", "", "docliId to delete")
	removeCommand.Parse(os.Args[2:])
	if removeCommand.Parsed() {
		if *docliId == "" {
			fmt.Println("Please supply docli name")
			return
		}
		strDocliId := *docliId
		err := api.DeleteDocli(strDocliId)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error deleting docli: " + strDocliId)
		}
	}
}