package common

import (
	"testing"
)

func BenchmarkMD5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MD5("")
	}
}
