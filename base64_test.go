package common

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestBase64(t *testing.T) {
	for i := 0; i < 10; i++ {
		input := strings.Repeat("1", i)
		if Base64(input) != base64.StdEncoding.EncodeToString([]byte(input)) {
			t.Fatal()
		}
	}
}

func BenchmarkRawBase64Short(b *testing.B) {
	input := "123456789"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString([]byte(input))
	}
}

func BenchmarkRawBase64Long(b *testing.B) {
	input := strings.Repeat("123456789", 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString([]byte(input))
	}
}

func BenchmarkBase64Short(b *testing.B) {
	input := "123456789"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64(input)
	}
}

func BenchmarkBase64Long(b *testing.B) {
	input := strings.Repeat("123456789", 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Base64(input)
	}
}
