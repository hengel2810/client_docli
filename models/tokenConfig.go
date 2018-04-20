package models

import "time"

type TokenConfig struct {
	AccessToken   string      `json:"access_token"`
	RefreshToken   string      `json:"refresh_token"`
	IdToken   string      `json:"id_token"`
	ExpiresIn   int      `json:"expires_in"`
	TokenType   string      `json:"token_type"`
	UserId		string		`json:"user_id"`
	ExpiringDate time.Time `json:"expiring_date"`
}
