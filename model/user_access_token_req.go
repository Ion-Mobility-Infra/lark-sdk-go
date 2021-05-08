package model

type UserAccessTokenReq struct {
	Code           string `json:"code"`
	AppAccessToken string `json:"app_access_token"`
	GrantType      string `json:"grant_type"`
}
