package models

type TokenResponse struct {
	User         User    `json:"user"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
}
