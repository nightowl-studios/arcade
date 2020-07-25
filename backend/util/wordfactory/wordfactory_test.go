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

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestGenerateWordList(t *testing.T) {
	expectedBehaviour := []string{"hello", "world", "alain"}
	actualBehaviour := generateWordList(3, annonymousFunc())
	for i := 0; i < 3; i++ {
		assertEqual(t, expectedBehaviour[i], actualBehaviour[i])
	}
}

func annonymousFunc() func(string) string {
	n := 0
	return func(string) string {
		testList := []string{"hello", "hello", "world", "world", "world", "alain"}
		n++
		return testList[n]
	}
}
