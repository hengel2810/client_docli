package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: cli <command> [<args>]")
		fmt.Println("The most commonly used git commands are: ")
		fmt.Println(" login   Login")
		fmt.Println(" logout  Logout")
		fmt.Println(" start  Upload image with YAML file to docli server")
		return
	}
	switch os.Args[1] {
	case "start":
		HandleUploadFromConfig()
	case "login":
		HandleLogin()
	case "logout":
		HandleLogout()
	case "list":
		HandleListDoclis()
	case "remove":
		HandleRemoveDocli()
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}