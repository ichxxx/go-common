package common

import (
	"math/rand"
	"testing"
)

func genTestSliceAndTarget(length int) ([]string, string) {
	s := make([]string, length)
	for i := 0; i < length; i++ {
		s[i] = ShortUUID()
	}
	target := ShortUUID()
	targetPos := rand.Intn(length)
	s[targetPos] = target
	return s, target
}

func BenchmarkInSlice(b *testing.B) {
	s, target := genTestSliceAndTarget(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InSlice(s, target)
	}
}
