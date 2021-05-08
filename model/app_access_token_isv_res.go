package model

type AppAccessTokenIsvRes struct {
	Code           int    `json:"code"`
	Expire         int    `json:"expire"`
	AppAccessToken string `json:"app_access_token"`
}
