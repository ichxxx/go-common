package common

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	if len(s) == 0 {
		return ""
	}
	md5Hex := md5.Sum(UnsafeBytes(s))
	return hex.EncodeToString(md5Hex[:])
}
