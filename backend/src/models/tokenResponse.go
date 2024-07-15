package models

type TokenResponse struct {
	User         User    `json:"user"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
}

type TokenAdminResponse struct {
	Admin         Admin    `json:"admin"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
}
