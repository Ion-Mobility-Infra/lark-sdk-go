package model

type GroupListBotRes struct {
	Code int `json:"code"`
	Data struct {
		Groups []struct {
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			Description string `json:"description"`
			ChatID      string `json:"chat_id"`
			OwnerOpenID string `json:"owner_open_id"`
			OwnerUserID string `json:"owner_user_id"`
		} `json:"groups"`
		HasMore   bool   `json:"has_more"`
		PageToken string `json:"page_token"`
	} `json:"data"`
	Msg string `json:"msg"`
}
