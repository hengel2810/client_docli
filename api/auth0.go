package api

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"sync"
	"github.com/skratchdot/open-golang/open"
	"github.com/hengel2810/client_docli/models"
	"encoding/json"
	"strings"
	"github.com/hengel2810/client_docli/config"
	"net/http"
	"io/ioutil"
)

type RequestBody struct {
	GrantType string `json:"grant_type"`
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TokenRequestBody
	RefreshRequestBody
}

type TokenRequestBody struct {
	Code string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
	IdToken string `json:"id_token"`
	TokenType string `json:"token_type"`
}

var wg sync.WaitGroup
func StartLoginProcess() {
	wg.Add(1)
	url := "https://hengel28.auth0.com/authorize?"
	audience := "audience=https://api.docli.com&"
	scope := "scope=offline_access&"
	response := "response_type=code&"
	clientId := "client_id=umW9qQGfeynUMuEZzino0IvF4d0U4QNs&"
	redirectURI := "redirect_uri=http://localhost&"
	url = url + audience + scope + response + clientId + redirectURI
	startLocalLoginServer()
	open.RunWith(url, "Google Chrome")
	wg.Wait()
}

func RequestToken(code string) {
	tokenRequestBody := RequestBody{
		GrantType:"authorization_code",
		ClientId:"umW9qQGfeynUMuEZzino0IvF4d0U4QNs",
		ClientSecret:"Dzp_mbLAEiYQdWoUDjdYyO1t0UVQrQGqXHQZ6XS941OOZnn-s69rhe-rqKzmF5Xe",
		TokenRequestBody: TokenRequestBody{
			Code: code,
			RedirectURI: "http://localhost",
		},
	}
	body, err := Auth0Request(tokenRequestBody)
	tokenConfig := &models.TokenConfig{}
	err = json.Unmarshal([]byte(string(body)), tokenConfig)
	if err !=  nil {
		fmt.Println(err)
		return
	}
	userId := userFromToken(tokenConfig.AccessToken)
	pipePos := strings.Index(userId, "|") + 1
	substringUserId := userId[pipePos:len(userId)]
	tokenConfig.UserId = substringUserId
	config.SaveTokenConfig(*tokenConfig)
}

func RefreshToken() (models.TokenConfig, error) {
	cfg, err := config.LoadTokenConfig()
	if err != nil {
		fmt.Println(err)
		return models.TokenConfig{}, err
	}
	tokenRequestBody := RequestBody{
		GrantType:"refresh_token",
		ClientId:"umW9qQGfeynUMuEZzino0IvF4d0U4QNs",
		ClientSecret:"Dzp_mbLAEiYQdWoUDjdYyO1t0UVQrQGqXHQZ6XS941OOZnn-s69rhe-rqKzmF5Xe",
		RefreshRequestBody: RefreshRequestBody{
			RefreshToken:cfg.RefreshToken,
		},
	}
	body, err := Auth0Request(tokenRequestBody)
	if err != nil{
		fmt.Println(err)
		fmt.Println("ERROR REFRESH")
		return models.TokenConfig{}, err
	}
	refreshResponse := &RefreshResponse{}
	err = json.Unmarshal([]byte(string(body)), refreshResponse)
	if err !=  nil {
		fmt.Println(err)
		return models.TokenConfig{}, err
	}
	cfg.AccessToken = refreshResponse.AccessToken
	config.SaveTokenConfig(cfg)
	return cfg, nil
}

func Auth0Request(requestBody RequestBody) ([]byte, error) {
	url := "https://hengel28.auth0.com/oauth/token"
	byteBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	payload := strings.NewReader(string(byteBody))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	if res.StatusCode == 200 {
		return body, nil
	} else {
		fmt.Print("wrong request status ")
		fmt.Print(res.StatusCode)
		return []byte{}, err
	}
}

func userFromToken(tokenstring string) string {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if token != nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil {
			return ""
		}
		userId := claims["sub"].(string)
		return userId
	}
	return ""
}