package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/random"
	"github.com/levigross/grequests"
	"github.com/quan-xiansheng/go-api-sdk/Response"
	"github.com/quan-xiansheng/go-api-sdk/utils"
	"log"
)

func (api *ApiClient) GetUsername(id int) (*Response.Response, error) {

	url := fmt.Sprintf("http://localhost:9001/api/sdk/user/name/%d", id)

	tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
	accessKey := api.AccessKey
	secretKey := api.SecretKey
	body := fmt.Sprintf("%d", id)
	headMap := map[string]string{
		"accessKey": accessKey,
		"nonce":     random.RandNumeral(3),
		"body":      body,
		"timestamp": tm.ToFormat(),
		"sign":      utils.GetSign(body, secretKey),
	}

	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: headMap,
	})
	if err != nil {
		return nil, err
	}

	var jsonResp Response.Response
	if err := json.Unmarshal(resp.Bytes(), &jsonResp); err != nil {
		log.Println("解析 JSON 出错:", err.Error())
		return nil, errors.New("服务繁忙，请稍后再试~~")
	}
	return &jsonResp, nil
}

func (api *ApiClient) GetRandomNum() string {

	url := fmt.Sprintf("http://localhost:8888/api/randomNum")

	tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
	accessKey := api.AccessKey
	secretKey := api.SecretKey
	body := fmt.Sprintf("%d", 1)
	headMap := map[string]string{
		"accessKey": accessKey,
		"nonce":     random.RandNumeral(3),
		"body":      body,
		"timestamp": tm.ToFormat(),
		"sign":      utils.GetSign(body, secretKey),
	}

	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: headMap,
	})
	if err != nil {
		return ""
	}
	return resp.String()

}
