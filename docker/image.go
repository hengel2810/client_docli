package docker

import (
	"github.com/docker/docker/client"
	"fmt"
	"io/ioutil"
	"golang.org/x/net/context"
)

func CopyImage(imageId, path string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return false
	}
	imageIds := []string{imageId}
	ioReadClose, err := cli.ImageSave(context.Background(), imageIds);
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer ioReadClose.Close()
	content, err :=  ioutil.ReadAll(ioReadClose)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		err := ioutil.WriteFile(path, content, 0644)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
