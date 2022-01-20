package common

import (
	"github.com/google/uuid"
)

func init()  {
	uuid.EnableRandPool()
}

var uuidChars = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func UUID() string {
	return uuid.NewString()
}

func ShortUUID() string {
	long := UUID()
	short := make([]byte, 8)
	buf := make([]byte, 4)
	for i := 0; i < 8; i++ {
		buf = buf[:0]
		for j := i*4; j < 36 && len(buf) < 4; j++ {
			if long[j] != '-' {
				buf = append(buf, long[j])
			}
		}
		short[i] = uuidChars[parseHex(UnsafeString(buf)) % 0x3E]
	}
	return UnsafeString(short)
}

func parseHex(s string) int {
	var (
		n int
		d byte
	)
	for _, c := range []byte(s) {
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		case 'a' <= c && c <= 'z':
			d = c - 'a' + 10
		default:
			return 0
		}

		n *= 16
		n += int(d)
	}
	return n
}
