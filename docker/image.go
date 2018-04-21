package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"github.com/hengel2810/client_docli/models"
	"errors"
)

func UploadDockerImage(docliObject models.DocliConfigObject) error {
	err := tagImage(docliObject.OriginalName, docliObject.FullName)
	if err != nil {
		return errors.New("Error while tagging image")
	}
	err = pushImage(docliObject.FullName)
	if err != nil {
		return errors.New("Error while pushing image")
	}
	err = removeTaggedImage(docliObject.FullName)
	if err != nil {
		return errors.New("Error while untagging image")
	}
	return nil
}

func tagImage(old string, new string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("Error while creating docker client")
	}
	err = cli.ImageTag(context.Background(), old, new)
	if err != nil {
		return errors.New("Error tagging image")
	}
	return nil
}

func pushImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("Error while creating docker client")
	}
	closer, err := cli.ImagePush(context.Background(), image, types.ImagePushOptions{All: false, RegistryAuth:"123"})
	if err != nil {
		return errors.New("Error pushing image")
	}
	_, err = io.Copy(os.Stdout, closer)
	if err != nil {
		return errors.New("Error copy image")
	}
	closer.Close()
	return nil
}

func removeTaggedImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.New("Error while creating docker client")
	}
	_, err = cli.ImageRemove(context.Background(), image, types.ImageRemoveOptions{Force:true})
	if err != nil {
		return errors.New("Error removing image")
	}
	return nil
}