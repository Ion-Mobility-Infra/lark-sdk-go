package lark

import (
	"bytes"
	"encoding/json"
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

// URLS for user access token
const POST_USER_ACCESS_TOKEN = "/authen/v1/access_token"

// Group List for Bot
const GET_GROUP_LIST = "/chat/v4/list"

// Client represents a client instance of app configuration
// that interacts with Lark Open APIs
type Client struct {
	appID             string
	appSecret         string
	appTicket         string
	internal          bool
	appAccessToken    string // set after an ObtainAppAccessToken call is made
	code              int    // set after an ObtainAppAccessToken call is made
	tenantAccessToken string // set after an ObtainTenantAccessTokeb call is made
}

func NewClient(internal bool) *Client {
	appID := os.Getenv("LARK_APP_ID")
	appSecret := os.Getenv("LARK_APP_SECRET")
	var appTicket string
	if internal == false {
		// an ISV app requires App Ticket
		appTicket = os.Getenv("LARK_APP_TICKET")
	}
	return &Client{
		appID:     appID,
		appSecret: appSecret,
		appTicket: appTicket, // only applicable for ISV apps. Not internal apps.
		internal:  internal,
	}
}

func (c *Client) ObtainAppAccessToken() *Client {

	// prepare data and url
	req := structToMap(model.AppAccessTokenInternalReq{
		AppID:     c.appID,
		AppSecret: c.appSecret,
	})
	url := LARKSUITE_OPEN_API + POST_APP_ACCESS_TOKEN_INTERNAL
	if c.internal == false {
		req = structToMap(model.AppAccessTokenIsvReq{
			AppID:     c.appID,
			AppSecret: c.appSecret,
			AppTicket: c.appTicket,
		})
		url = LARKSUITE_OPEN_API + POST_APP_ACCESS_TOKEN
	}

	body := post(req, url)

	// deserialize
	if c.internal {
		var res model.AppAccessTokenInternalRes
		json.Unmarshal(body, &res)
		c.appAccessToken = res.AppAccessToken
		c.code = res.Code
	} else {
		var res model.AppAccessTokenIsvRes
		json.Unmarshal(body, &res)
		c.appAccessToken = res.AppAccessToken
		c.code = res.Code
	}

	return c
}

func (c *Client) ObtainTenantAccessToken() *Client {
	req := structToMap(model.TenantAccessTokenReq{
		AppID:     c.appID,
		AppSecret: c.appSecret,
	})
	url := LARKSUITE_OPEN_API + POST_TENANT_ACCESS_TOKEN_INTERNAL

	body := post(req, url)

	var res model.TenantAccessTokenRes
	json.Unmarshal(body, &res)
	c.tenantAccessToken = res.TenantAccessToken
	c.code = res.Code
	return c
}

func (c *Client) ObtainBotGroupList() *model.GroupListBotRes {
	url := LARKSUITE_OPEN_API + GET_GROUP_LIST
	body := get(url, c.appAccessToken)

	var res model.GroupListBotRes
	json.Unmarshal(body, &res)
	return &res
}

func (c *Client) ObtainUserAccessToken(code string) *model.UserAccessTokenRes {
	req := structToMap(model.UserAccessTokenReq{
		Code:           code,
		AppAccessToken: c.appAccessToken,
		GrantType:      "authorization_code",
	})
	url := LARKSUITE_OPEN_API + POST_USER_ACCESS_TOKEN

	body := post(req, url)

	var res model.UserAccessTokenRes
	json.Unmarshal(body, &res)
	return &res
}

func post(req interface{}, url string) (body []byte) {
	// serialize and make API call
	jsonData, _ := json.Marshal(req)
	jsonBytes := bytes.NewBuffer(jsonData)

	resp, err := http.Post(url, "application/json", jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// response
	body, _ = ioutil.ReadAll(resp.Body)
	return body
}

func get(url string, accessToken string) (body []byte) {
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + accessToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// response
	body, _ = ioutil.ReadAll(resp.Body)
	return body
}
