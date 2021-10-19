package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	md5Hex := md5.Sum([]byte(strings.Join(s, "")))
	return hex.EncodeToString(md5Hex[:])
}