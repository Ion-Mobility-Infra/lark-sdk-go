package model

type AppAccessTokenInternalReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}
