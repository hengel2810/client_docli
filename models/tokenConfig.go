package models

type TokenConfig struct {
	AccessToken   string      `json:"access_token"`
	IdToken   string      `json:"id_token"`
	ExpiresIn   int      `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	UserId		string		`json:"user_id"`
}
