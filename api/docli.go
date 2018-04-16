package api

import (
	"bytes"
	"net/http"
	"fmt"
	"github.com/hengel2810/client_docli/config"
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"encoding/json"
	"github.com/hengel2810/client_docli/controller"
)

func PostImageData(uploadImage models.DocliObject) {
	if controller.DocliObjectValid(uploadImage) == false {
		return
	}
	//url := "http://46.101.222.225:8000/image"
	url := "http://localhost:8000/image"
	data, err := json.Marshal(uploadImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("error create NewRequest")
		return
	}
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		fmt.Println("error load config")
		return
	}
	req.Header.Set("Authorization", "Bearer " + cfg.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}