package utils

import "github.com/duke-git/lancet/v2/cryptor"

func GetSign(body string, secreKey string) string {

	secre := cryptor.Sha1(body + secreKey)
	return secre
}
