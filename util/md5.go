package util

import (
	"crypto/md5"
	"encoding/base64"
)

// Md5
func Md5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
