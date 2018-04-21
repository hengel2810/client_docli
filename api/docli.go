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
)

func PostImageData(uploadImage models.DocliConfigObject) error {
	if controller.DocliObjectValid(uploadImage) == false {
		return errors.New("invalid docli object")
	}
	url := "https://api.valas.cloud/docli"
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
		cfg, err = RefreshToken()
		if err != nil {
			return err
		}
	//}
	req.Header.Set("Authorization", "Bearer " + cfg.AccessToken)
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

func GetDoclis(userId string) ([]models.DocliObject, error) {
	//url := "https://api.valas.cloud/images"
	url := "http://localhost:8000/doclis?userId=" + userId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		return []models.DocliObject{}, errors.New("error load config")
	}
	//if cfg.ExpiringDate.Local().Second() < time.Now().Local().Second() {
		cfg, err = RefreshToken()
		if err != nil {
			return []models.DocliObject{},err
		}
	//}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []models.DocliObject{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []models.DocliObject{}, err
	}
	if res.StatusCode == 200 {
		var arrDocli []models.DocliObject
		err = json.Unmarshal(body, &arrDocli)
		if err != nil {
			return []models.DocliObject{}, err
		}
		return arrDocli, nil
	} else {
		return []models.DocliObject{}, errors.New("wrong status" + strconv.Itoa(res.StatusCode))
	}
}

func DeleteDocli(docliId string) error {
	//url := "https://api.valas.cloud/images"
	url := "http://localhost:8000/docli?docliId=" + docliId
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		return errors.New("error load config")
	}
	//if cfg.ExpiringDate.Local().Second() < time.Now().Local().Second() {
		cfg, err = RefreshToken()
		if err != nil {
			return err
		}
	//}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	} else {
		return errors.New("wrong status" + strconv.Itoa(res.StatusCode))
	}
}