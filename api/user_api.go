package api

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/random"
	"github.com/levigross/grequests"
	"github.com/quanxiangsheng/go-api-sdk/utils"
)

func (api *ApiClient) GetUsername(id int) string {

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
		return ""
	}
	return resp.String()

}
