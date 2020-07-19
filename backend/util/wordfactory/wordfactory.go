package wordfactory

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bseto/arcade/backend/log"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

var (
	Dir  string = filepath.Join("util", "wordfactory")
	File string = "wordbank.txt"
)

func WordGenerator(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Errorf("There is an error in reading file.", err)
		return "", err
	}

	allWords := string(b[:])
	wordList := []string{allWords}

	words := strings.Split(wordList[0], "\r\n")
	pickWord := randInt(0, len(wordList))
	return words[pickWord], nil
}

func WordGenerator2(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	fileInfo, err := file.Stat()
	if err != nil {
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
