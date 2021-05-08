package model

type SendTextRes struct {
	Code int `json:"code"`
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
	Msg string `json:"msg"`
}
