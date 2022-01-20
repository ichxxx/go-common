package common

import (
	"github.com/spf13/cast"
)

func Int(i interface{}) int {
	return cast.ToInt(i)
}
