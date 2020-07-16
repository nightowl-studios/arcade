package wordfactory

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func randInt2(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
func BenchmarkWordGenerator1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordGenerator()
	}
}

func BenchmarkWordGenerator2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordGenerator2()
	}
}

// func TestWordFactoryFunction(t *testing.T) {
// 	ans, _ := testWordGenerator2(0)
// 	t.Error(ans)
// }

func TestYourFunction(t *testing.T) {
	file, err := os.Open("wordbank.txt")
	if err != nil {
		fmt.Println("file could not be read", err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("file info cannot be read", err)
	}

	var fileSize int = int(fileInfo.Size())
	randomLocation := randInt2(0, fileSize-10)
	for randomLocation <= 10 {
		randomLocation = randInt2(0, fileSize-10)
	}

	reader := bufio.NewReader(file)
	reader.Discard(randomLocation)

	// data, _, _ := reader.ReadLine()
	// data, _, _ = reader.ReadLine()
	// t.Error(string(data))
	file.Close()
}
