package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/model"
)

const LARKSUITE_OPEN_API = "https://open.larksuite.com/open-apis"

// URLs for API Access Tokens
const POST_APP_ACCESS_TOKEN = "/auth/v3/app_access_token"
const POST_APP_ACCESS_TOKEN_INTERNAL = "/auth/v3/app_access_token/internal"
const POST_TENANT_ACCESS_TOKEN = "/auth/v3/tenant_access_token"
const POST_TENANT_ACCESS_TOKEN_INTERNAL = "/auth/v3/tenant_access_token/internal"

type Client struct {
	appID          string
	appSecret      string
	appAccessToken string
	internal       bool
}

func NewClient(internal bool) *Client {
	larkAppID := os.Getenv("LARK_APP_ID")
	larkAppSecret := os.Getenv("LARK_APP_SECRET")
	return &Client{
		appID:     larkAppID,
		appSecret: larkAppSecret,
		internal:  internal,
	}
}

func (c *Client) ObtainAppAccessToken() *Client {
	req := structToMap(model.AppAccessTokenInternalReq{
		AppID:     c.appID,
		AppSecret: c.appSecret,
	})

	jsonData, _ := json.Marshal(req)
	jsonBytes := bytes.NewBuffer(jsonData)
	url := LARKSUITE_OPEN_API + POST_APP_ACCESS_TOKEN_INTERNAL
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var res model.AppAccessTokenInternalRes
	json.Unmarshal(body, &res)

	c.appAccessToken = res.AppAccessToken
	return c
}
