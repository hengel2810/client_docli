package api

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"fmt"
	"github.com/hengel2810/client_docli/config"
	"io/ioutil"
	"github.com/hengel2810/client_docli/models"
	"encoding/json"
)

func PostImageData(uploadImage models.DockerImageUpload) {
	//url := "http://46.101.222.225:8000/image"
	url := "http://localhost:8000/image"
	data, err := json.Marshal(uploadImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("error create NewRequest")
		return
	}
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		fmt.Println("error load config")
		return
	}
	req.Header.Set("Authorization", "Bearer " + cfg.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func PostFile(filename string, filepath string, targetUrl string, imageId string) (error, int) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.WriteField("image", imageId)
	if err != nil {
		fmt.Println("error writing image field")
		return err, 900
	}
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
