package test

import (
	"fmt"
	quanSdk "github.com/quan-xiansheng/go-api-sdk"
	"testing"
)

var client = quanSdk.NewApiClient("ABC", "abc")

func TestGetUsername(t *testing.T) {

	id := 1
	username := client.GetUsername(id)
	fmt.Println(username)
}
