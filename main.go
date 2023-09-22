package apisdk

import (
	"github.com/quan-xiansheng/go-api-sdk/api"
)

func NewApiClient(AccessKey, SecretKey string) *api.ApiClient {
	return &api.ApiClient{
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}
}
