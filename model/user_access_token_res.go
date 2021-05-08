package model

type UserAccessTokenRes struct {
	Name         string `json:"name"`
	OpenID       string `json:"open_id"`
	ExpiresIn    int    `json:"expires_in"`
	TenantKey    string `json:"tenant_key"`
	EnName       string `json:"en_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	AvatarURL    string `json:"avatar_url"`
}
