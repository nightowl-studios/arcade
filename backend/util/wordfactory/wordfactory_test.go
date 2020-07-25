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

func TestGenerateWordList(t *testing.T) {
	wordList := generateWordList(3, wordGenerator())
	duplicateCheck := make(map[string]bool)
	for _, word := range wordList {
		duplicateCheck[word] = true
	}

	if len(duplicateCheck) != len(wordList) {
		t.Errorf(
			"there were :%v duplicates inside :%v",
			len(wordList)-len(duplicateCheck),
			wordList,
		)
	}
}

func wordGenerator() func(string) string {
	n := 0
	testList := []string{"hello", "hello", "world", "world", "world", "alain"}
	return func(string) string {
		n++
		return testList[n]
	}
}
