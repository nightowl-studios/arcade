package wordfactory

import (
	"math/rand"
	"testing"
	"time"
)

func randInt2(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
func BenchmarkWordGenerator1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordGenerator(File)
	}
}

func BenchmarkWordGenerator2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordGenerator2(File)
	}
}
