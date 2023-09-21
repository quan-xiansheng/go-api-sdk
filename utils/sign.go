package utils

import "github.com/duke-git/lancet/cryptor"

func GetSign(body string, secreKey string) string {

	secre := cryptor.Sha1(body + secreKey)
	return secre
}
