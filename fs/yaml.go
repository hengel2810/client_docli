package fs

import (
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"github.com/ghodss/yaml"
)

func ReadConfig(filePath string) (models.DocliObject, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return models.DocliObject{}, err
	}
	var docli models.DocliObject
	err = yaml.Unmarshal(yamlFile, &docli)
	if err != nil {
		return models.DocliObject{}, err
	}
	return docli, nil
}
