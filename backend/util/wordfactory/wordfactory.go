package nullable

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func wordGenerator() {
	b, err := ioutil.ReadFile("wordbank.txt")

	if err != nil {
		fmt.Println("There is an error in reading file.", err)
	}

	allWords := string(b[:])
	wordList := []string{allWords}

	words := strings.Split(wordList[0], ",")
	pickWord := rand.Intn(len(words))
	fmt.Println(words[pickWord])
}
