package common

import (
	"testing"
	"time"
)

func BenchmarkRetry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Retry(func() error {
			return nil
		}).Do()
	}
}

func BenchmarkRetryWithTimes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Retry(func() error {
			return nil
		}).Times(5).Do()
	}
}

func BenchmarkRetryWithWait(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Retry(func() error {
			return nil
		}).Wait(time.Millisecond).Do()
	}
}
