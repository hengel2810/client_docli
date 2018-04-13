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
		fmt.Println(" upload  Upload image to docli server")
		return
	}
	switch os.Args[1] {
	case "upload":
		HandleUpload()
	case "login":
		HandleLogin()
	case "logout":
		HandleLogout()
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}