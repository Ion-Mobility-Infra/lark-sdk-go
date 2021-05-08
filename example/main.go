package main

import (
	"fmt"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/lark"
)

func main() {
	fmt.Println("Example API call to LarkSuite API")

	result := lark.NewClient(true).
		ObtainAppAccessToken().
		ObtainTenantAccessToken().
		ObtainBotGroupList()

	var chatIDs []string
	if len(result.Data.Groups) > 0 {
		for _, group := range result.Data.Groups {
			chatIDs = append(chatIDs, group.ChatID)
		}
	}

	fmt.Println(chatIDs)

}
