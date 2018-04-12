package config

import (
	"encoding/json"
	"github.com/hengel2810/client_docli/models"
	"io/ioutil"
	"os"
	"fmt"
)

func SaveTokenConfig(tokenConfig models.TokenConfig) error {
	bytes, err := json.MarshalIndent(tokenConfig, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("config.json", bytes, 0644)
}

func LoadTokenConfig() (models.TokenConfig, error) {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		return models.TokenConfig{}, err
	}
	var tokenConfig models.TokenConfig
	err = json.Unmarshal(bytes, &tokenConfig)
	if err != nil {
		return models.TokenConfig{}, err
	}
	return tokenConfig, nil
}

func RemoveTokenConfig() {
	err := os.Remove("config.json")
	if err != nil {
		fmt.Println(err)
	}
}

func ConfigValid() bool {
	tokenConfig, err := LoadTokenConfig()
	if err != nil {
		return false
	}
	if tokenConfig.AccessToken == "" || tokenConfig.IdToken == "" || tokenConfig.TokenType == "" || tokenConfig.ExpiresIn == 0 {
		return false
	}
	return true
}