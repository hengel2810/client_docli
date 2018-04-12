package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"io"
	"bytes"
	"mime/multipart"
	"github.com/hengel2810/client_docli/login"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: cli <command> [<args>]")
		fmt.Println("The most commonly used git commands are: ")
		fmt.Println(" login   Login")
		fmt.Println(" logout  Logout")
		fmt.Println(" upload  Upload image to docli server")
		return
	}
	switch os.Args[1] {
	case "upload":
		handleUpload()
	case "login":
		login.HandleLogin()
	case "logout":
		login.HandleLogout()
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}

func handleUpload() {
	uploadCommand := flag.NewFlagSet("upload", flag.ExitOnError)
	imageId := uploadCommand.String("image", "", "image id or name")
	uploadCommand.Parse(os.Args[2:])
	if uploadCommand.Parsed() {
		if *imageId == "" {
			fmt.Println("Please supply image id or name")
			return
		}
		tempPath := os.TempDir()
		fileName := *imageId + ".tar"
		filePath := tempPath + fileName
		copied := copyImage(*imageId, filePath)
		if copied {
			postError, statusCode := postFile(fileName, filePath, "http://localhost:8000/image", *imageId)
			if postError == nil && statusCode == 200 {
				fmt.Println("Image uploaded")
			} else {
				fmt.Println("Error while uploading image")
			}
		} else {
			fmt.Println("Error while copying image")
		}
	}
}

func copyImage(imageId, path string) bool {
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

func postFile(filename string, filepath string, targetUrl string, imageId string) (error, int) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err, 900
	}
	fh, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file")
		return err, 900
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err, 900
	}
	err = bodyWriter.WriteField("image", imageId)
	if err != nil {
		fmt.Println("error writing image field")
		return err, 900
	}
	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		fmt.Println("error closing bodywriter")
		return err, 900
	}
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err, 900
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, 900
	}
	return nil, resp.StatusCode
}