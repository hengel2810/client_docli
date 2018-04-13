package api

import (
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
)

func PostFile(filename string, filepath string, targetUrl string, imageId string) (error, int) {
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
