package common

import (
	"encoding/base64"
)

func Base64(s string) string {
	length := len(s)/3
	if length*3 != len(s) {
		length++
	}
	buf := make([]byte, length*4)
	base64.StdEncoding.Encode(buf, UnsafeBytes(s))
	return UnsafeString(buf)
}
