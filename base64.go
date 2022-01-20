package common

import (
	"encoding/base64"
)

func Base64(s string) string {
	x := len(s)/3
	if x*3 != len(s) {
		x++
	}
	buf := make([]byte, x*4)
	base64.StdEncoding.Encode(buf, UnsafeBytes(s))
	return UnsafeString(buf)
}

func DecodeBase64(s string) string {
	blank := 0
	for i := len(s)-1; i >= 0 && s[i] == '='; i-- {
		blank++
	}
	buf := make([]byte, (len(s)/4)*3-blank)
	_, _ = base64.StdEncoding.Decode(buf, UnsafeBytes(s))
	return UnsafeString(buf)
}
