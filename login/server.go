package login

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"github.com/hengel2810/client_docli/config"
	"encoding/json"
	"github.com/hengel2810/client_docli/models"
	"sync"
	"github.com/skratchdot/open-golang/open"
	"time"
)

var srv *http.Server
var wg sync.WaitGroup

func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":3000"}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/code", codeHandler)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	requestToken(code)
	htmlContent, err := ioutil.ReadFile("static/loggedin.html")
	if err == nil {
		w.Write(htmlContent)
	} else {
		fmt.Println(err)
	}
	go stopServer()
}

func stopServer() {
	time.Sleep(1 * time.Second)
	defer wg.Done()
}

func StartLoginServer() {
	wg.Add(1)
	srv = startHttpServer()
	open.Run("http://localhost:3000/login.html")
	wg.Wait()
}

func requestToken(code string) {
	fmt.Println(code)
	url := "https://hengel28.auth0.com/oauth/token"
	payload := strings.NewReader("{\n  \"grant_type\": \"authorization_code\",\n  \"client_id\": \"umW9qQGfeynUMuEZzino0IvF4d0U4QNs\",\n  \"client_secret\": \"Dzp_mbLAEiYQdWoUDjdYyO1t0UVQrQGqXHQZ6XS941OOZnn-s69rhe-rqKzmF5Xe\",\n  \"code\": \"" + code + "\",\n  \"redirect_uri\": \"http://localhost:3000/login.html\"\n}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode == 200 {
		tokenConfig := &models.TokenConfig{}
		err := json.Unmarshal([]byte(string(body)), tokenConfig)
		if err !=  nil {
			fmt.Println(err)
		}
		config.SaveTokenConfig(*tokenConfig)
	} else {
		fmt.Println("TOKEN ERROR")
	}
	//defer wg.Done()
}
