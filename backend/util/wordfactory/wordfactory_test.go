package wordfactory

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

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
	randomLocation := rand.Intn(fileSize)
	t.Error(fileSize)
	t.Error(randomLocation)
	reader := bufio.NewReader(file)

	reader.Discard(randomLocation)

	data, _, _ := reader.ReadLine()
	data, _, _ = reader.ReadLine()
	t.Error(string(data))
	file.Close()
}
