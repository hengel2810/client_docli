package api

import (
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/http"
	"fmt"
	"github.com/hengel2810/client_docli/config"
)

func PostFile(filename string, filepath string, targetUrl string, imageId string) (error, int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error open file")
		return err, 900
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error create form file")
		return err, 900
	}
	err = writer.WriteField("image", imageId)
	if err != nil {
		fmt.Println("error writing image field")
		return err, 900
	}
	io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		fmt.Println("error writer close")
		return err, 900
	}
	r, err := http.NewRequest("POST", targetUrl, body)
	if err != nil {
		fmt.Println("error creating reuqest")
		return err, 900
	}
	r.Header.Add("Content-Type", writer.FormDataContentType())
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		fmt.Println("error load config")
		return err, 900
	}
	r.Header.Add("Authorization", "Bearer " + cfg.AccessToken)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	return nil, res.StatusCode
}
