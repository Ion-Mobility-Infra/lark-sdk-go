package model

type SendTextReq struct {
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
	Email   string `json:"email"`
	ChatID  string `json:"chat_id"`
	OpenID  string `json:"open_id"`
	RootID  string `json:"root_id"`
	UserID  string `json:"user_id"`
	MsgType string `json:"msg_type"`
}
