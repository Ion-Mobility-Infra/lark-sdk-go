package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/api"
	"github.com/Ion-Mobility-Infra/lark-sdk-go/model"
)

func main() {
	fmt.Println("Example API call to LarkSuite API")
	larkAppID := os.Getenv("LARK_APP_ID")
	larkAppSecret := os.Getenv("LARK_APP_SECRET")
	req := model.AppAccessTokenInternalReq{
		AppID:     larkAppID,
		AppSecret: larkAppSecret,
	}
	m := structToMap(req)
	fmt.Println(m)
	// client := &http.Client{}
	// data := model.AppAccessTokenInternalReq{}
	jsonData, _ := json.Marshal(m)
	jsonBytes := bytes.NewBuffer(jsonData)
	url := api.LARKSUITE_OPEN_API + api.POST_APP_ACCESS_TOKEN_INTERNAL
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var b model.AppAccessTokenInternalRes
	json.Unmarshal(body, &b)
	fmt.Println(b.AppAccessToken)
	fmt.Println(b.Code)
	fmt.Println(b.Expire)

}

func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
