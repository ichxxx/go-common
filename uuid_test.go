package common

import (
	"testing"
)

func BenchmarkUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func BenchmarkShortUUID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ShortUUID()
	}
}
