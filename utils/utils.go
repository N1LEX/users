package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Hash(value string) string {
	p := []byte(value)
	h := sha256.New()
	h.Write(p)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
