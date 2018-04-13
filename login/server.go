package login

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"sync"
	"time"
	"github.com/hengel2810/client_docli/api"
	"github.com/skratchdot/open-golang/open"
)

var srv *http.Server
var wg sync.WaitGroup

func StartLoginServer() {
	wg.Add(1)
	srv = startHttpServer()
	url := "http://localhost:3000/login.html"
	open.Run(url)
	fmt.Println("Waiting for login on " + url +" ...")
	wg.Wait()
}

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
	api.RequestToken(code)
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
