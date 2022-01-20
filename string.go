package common

import (
	"strings"
	"unsafe"

	"github.com/spf13/cast"
)

func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String(i interface{}, defaultString ...string) string {
	s := cast.ToString(i)
	if len(s) == 0 && len(defaultString) > 0 {
		s = strings.Join(defaultString, "")
	}
	return s
}
