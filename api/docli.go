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
	"github.com/hengel2810/client_docli/login"
)

func PostImageData(uploadImage models.DocliObject) error {
	if controller.DocliObjectValid(uploadImage) == false {
		return errors.New("invalid docli object")
	}
	url := "https://api.valas.cloud/image"
	//url := "http://localhost:8000/image"
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
	//if cfg.ExpiringDate.Local().Second() < time.Now().Local().Second() {
		cfg, err = login.RefreshToken()
		if err != nil {
			return err
		}
	//}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("error doing request")
	}
	_, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return errors.New("error reading request body")
	} else {
		if resp.StatusCode == 200 {
			return nil
		} else {
			errorMsg := "wrong status " +  strconv.Itoa(resp.StatusCode)
			return errors.New(errorMsg)
		}
	}
}