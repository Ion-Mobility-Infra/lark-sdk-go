package model

type TenantAccessTokenRes struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	TenantAccessToken string `json:"tenant_access_token"`
}
