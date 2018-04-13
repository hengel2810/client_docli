package api

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"encoding/json"
	"github.com/hengel2810/client_docli/config"
)

func RequestToken(code string) {
	fmt.Println(code)
	url := "https://hengel28.auth0.com/oauth/token"
	payload := strings.NewReader("{\n  \"grant_type\": \"authorization_code\",\n  \"client_id\": \"umW9qQGfeynUMuEZzino0IvF4d0U4QNs\",\n  \"client_secret\": \"Dzp_mbLAEiYQdWoUDjdYyO1t0UVQrQGqXHQZ6XS941OOZnn-s69rhe-rqKzmF5Xe\",\n  \"code\": \"" + code + "\",\n  \"redirect_uri\": \"http://localhost:3000/login.html\"\n}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode == 200 {
		tokenConfig := &models.TokenConfig{}
		err := json.Unmarshal([]byte(string(body)), tokenConfig)
		if err !=  nil {
			fmt.Println(err)
		}
		config.SaveTokenConfig(*tokenConfig)
	} else {
		fmt.Println("TOKEN ERROR")
	}
}
