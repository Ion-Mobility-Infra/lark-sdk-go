package main

import (
	"fmt"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/lark"
)

func main() {
	fmt.Println("Example API calls to LarkSuite API")

	// Instantiate a lark API client and
	// obtain our app access token, and
	// obtain our tenant access token
	client := lark.NewClient(true).
		ObtainAppAccessToken().
		ObtainTenantAccessToken()

	// This API call to retrieve a list of groups that the bot belongs in
	// requires the app access token
	groupList := client.ObtainBotGroupList()

	var chatIDs []string
	if len(groupList.Data.Groups) > 0 {
		for _, group := range groupList.Data.Groups {
			chatIDs = append(chatIDs, group.ChatID)
		}
	}

	fmt.Println(chatIDs)

	// This API call to send a text message requires the tenant access token
	for _, chatID := range chatIDs {
		sendTextResp := client.SendTextMessage(chatID, "Hello there, this is a test message from R2-D2.\nI am built with an open source lark-go-sdk on https://github.com/Ion-Mobility-Infra/lark-sdk-go")
		fmt.Println(sendTextResp)
	}

}
