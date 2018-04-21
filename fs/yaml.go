package fs

import (
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"github.com/ghodss/yaml"
)

func ReadConfig(filePath string) (models.DocliConfigObject, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return models.DocliConfigObject{}, err
	}
	var docli models.DocliConfigObject
	err = yaml.Unmarshal(yamlFile, &docli)
	if err != nil {
		return models.DocliConfigObject{}, err
	}
	return docli, nil
}
