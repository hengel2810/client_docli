package docker

import (
	"github.com/docker/docker/client"
	"fmt"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"github.com/hengel2810/client_docli/config"
	"github.com/satori/go.uuid"
	"github.com/hengel2810/client_docli/models"
	"errors"
)

func UploadDockerImage(imageId string) (models.DockerImageUpload, error) {
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		return models.DockerImageUpload{}, errors.New("error load config")
	}
	registryURL := "46.101.222.225:5000"

	uniqueImageTag, err := uuid.NewV4()
	if err != nil {
		return models.DockerImageUpload{}, errors.New("error on unique tag creation")
	}
	newImageName := registryURL + "/" + cfg.UserId + "/" + uniqueImageTag.String() + "/" + imageId
	result := tagImage(imageId, newImageName)
	if result == false {
		return models.DockerImageUpload{}, errors.New("Error while tagging image")
	}
	result = pushImage(newImageName)
	if result == false {
		return models.DockerImageUpload{}, errors.New("Error while pushing image")
	}
	result = removeTaggedImage(newImageName)
	if result == false {
		fmt.Println("Error while untagging image")
		return models.DockerImageUpload{}, errors.New("Error while untagging image")
	}
	return models.DockerImageUpload{FullName:newImageName, UserId:cfg.UserId, OriginalName:imageId, UniqueId:uniqueImageTag.String()}, nil
}

func tagImage(old string, new string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = cli.ImageTag(context.Background(), old, new)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func pushImage(image string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return false
	}
	closer, err := cli.ImagePush(context.Background(), image, types.ImagePushOptions{All: false, RegistryAuth:"123"})
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = io.Copy(os.Stdout, closer)
	if err != nil {
		fmt.Println(err)
		return false
	}
	closer.Close()
	return true
}

func removeTaggedImage(image string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return false
	}
	resp, err := cli.ImageRemove(context.Background(), image, types.ImageRemoveOptions{Force:true})
	fmt.Println(resp)
	return true
}