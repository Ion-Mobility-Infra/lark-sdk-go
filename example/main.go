package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Example API call to LarkSuite API")
	larkAppID := os.Getenv("LARK_APP_ID")
	larkAppSecret := os.Getenv("LARK_APP_SECRET")
	fmt.Println(larkAppID)
	fmt.Println(larkAppSecret)
}
