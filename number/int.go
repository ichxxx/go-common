package number

import (
	"github.com/spf13/cast"
)

func ToInt(i interface{}) int {
	return cast.ToInt(i)
}