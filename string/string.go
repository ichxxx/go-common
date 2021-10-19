package string

import (
	"strings"
	"unsafe"

	"github.com/spf13/cast"
)

func Bytes(s string) []byte {
	tmp := (*[2]uintptr)(unsafe.Pointer(&s))
	x := [3]uintptr{tmp[0], tmp[1], tmp[1]}
	return *(*[]byte)(unsafe.Pointer(&x))
}

func From(i interface{}, _default ...string) string {
	s := cast.ToString(i)
	if len(s) == 0 && len(_default) > 0 {
		s = strings.Join(_default, "")
	}
	return s
}