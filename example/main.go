package main

import (
	"fmt"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/lark"
)

func main() {
	fmt.Println("Example API call to LarkSuite API")

	client := lark.NewClient(true).ObtainAppAccessToken()
	fmt.Println(client)

}
