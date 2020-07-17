package wordfactory

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func WordGenerator() string {
	b, err := ioutil.ReadFile("wordbank.txt")

	if err != nil {
		fmt.Println("There is an error in reading file.", err)
	}

	allWords := string(b[:])
	wordList := []string{allWords}

	words := strings.Split(wordList[0], "\r\n")
	pickWord := randInt(0, len(wordList))
	return words[pickWord]
}

func WordGenerator2() (string, error) {
	file, err := os.Open("wordbank.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("file could not be read", err)
		return "", err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("file info cannot be read", err)
		return "", err
	}

	var fileSize int = int(fileInfo.Size())
	randomLocation := randInt(0, fileSize-10)
	for randomLocation <= 10 {
		randomLocation = randInt(0, fileSize-10)
	}

	reader := bufio.NewReader(file)

	reader.Discard(randomLocation)

	data, _, _ := reader.ReadLine()
	data, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println("Cannot read next line", err)
		return "", err
	}

	return (string(data)), nil
}

func testWordGenerator2(testrand int) (int, error) {
	file, err := os.Open("wordbank.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("file could not be read", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("file info cannot be read", err)
	}

	var fileSize int = int(fileInfo.Size())
	for testrand <= 10 {
		testrand = randInt(0, fileSize-10)
	}
	return testrand, nil
}
