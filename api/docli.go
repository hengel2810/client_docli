package api

import (
	"bytes"
	"net/http"
	"github.com/hengel2810/client_docli/config"
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"encoding/json"
	"github.com/hengel2810/client_docli/controller"
	"errors"
	"strconv"
	"fmt"
)

func PostImageData(uploadImage models.DocliObject) error {
	if controller.DocliObjectValid(uploadImage) == false {
		return errors.New("invalid docli object")
	}
	//url := "http://46.101.222.225:8000/image"
	url := "http://localhost:8000/image"
	data, err := json.Marshal(uploadImage)
	if err != nil {
		return errors.New("error json marshal docli object")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return errors.New("error create NewRequest")
	}
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		return errors.New("error load config")
	}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("error doing request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return errors.New("error reading request body")
	} else {
		fmt.Println(resp.StatusCode)
		if resp.StatusCode == 200 {
			return nil
		} else {
			fmt.Println(string(body))
			errorMsg := "wrong status " +  strconv.Itoa(resp.StatusCode)
			return errors.New(errorMsg)
		}
	}
}