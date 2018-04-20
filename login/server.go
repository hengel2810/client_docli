package login

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"time"
)

var srv *http.Server


func startLocalLoginServer() {
	srv := &http.Server{Addr: ":80"}
	http.HandleFunc("/", codeHandler)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code != "" {
		RequestToken(code)
		htmlContent, err := ioutil.ReadFile("static/loggedin.html")
		if err == nil {
			w.Write(htmlContent)
		} else {
			fmt.Println(err)
		}
		go stopServer()
	}
}

func stopServer() {
	time.Sleep(1 * time.Second)
	defer wg.Done()
}
