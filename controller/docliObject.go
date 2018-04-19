package controller

import (
	"github.com/satori/go.uuid"
	"github.com/hengel2810/client_docli/config"
	"github.com/hengel2810/client_docli/models"
	"errors"
)

func DocliObjectValid(docliObject models.DocliObject) bool {
	if docliObject.FullName == "" {
		return false
	}
	if docliObject.UserId == "" {
		return false
	}
	if docliObject.OriginalName == "" {
		return false
	}
	if docliObject.UniqueId == "" {
		return false
	}
	return true
}

func SetDocliObjectData(docliObject models.DocliObject) (models.DocliObject, error) {
	registryURL := "registry.valas.cloud"
	uniqueImageTag, err := uuid.NewV4()
	if err != nil {
		return docliObject, errors.New("Error creating uniqueImageId")
	}
	docliObject.UniqueId = uniqueImageTag.String()
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		return docliObject, errors.New("Error while loading config")
	}
	newImageName := registryURL + "/" + cfg.UserId + "/" + uniqueImageTag.String() + "/" + docliObject.OriginalName
	docliObject.FullName = newImageName
	return docliObject, nil
}
