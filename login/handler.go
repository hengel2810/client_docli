package login

import (
	"github.com/hengel2810/client_docli/config"
	"fmt"
)

func HandleLogin()  {
	configValid := config.ConfigValid()
	if !configValid {
		StartLoginServer()
	} else {
		fmt.Println("Already logged in. Please use 'docli logout' to logout before re-login")
	}
}

func HandleLogout() {
	config.RemoveTokenConfig()
}