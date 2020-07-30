package wordhint

import (
	"math/rand"
	"time"

	"github.com/bseto/arcade/backend/log"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

type WordHint interface {
	GiveHint(word string) string
}

type wordHint struct {
	previousHintIndexes []int
}

func Get() *wordHint {
	rand.Seed(time.Now().UTC().UnixNano())
	return &wordHint{}
}

func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (w wordHint) GiveHint(word string) string {
	if word == "" {
		log.Errorf("Give hint cannot operate on an empty word")
		return ""
	}
	hint := ""
	hintIndex := randInt(1, len(word))
	found := Find(w.previousHintIndexes, hintIndex)
	if string(word[hintIndex]) == " " {
		found = true
	}
	for found {
		hintIndex = randInt(1, len(word))
		found = Find(w.previousHintIndexes, hintIndex)
		if string(word[hintIndex]) == " " {
			found = true
		}
	}
	w.previousHintIndexes = append(w.previousHintIndexes, hintIndex)
	for wordLength := 0; wordLength < len(word); wordLength++ {
		isFound := Find(w.previousHintIndexes, wordLength)
		if isFound {
			hint += (string(word[wordLength]) + " ")
		} else if wordLength == len(word)-1 && !(isFound) {
			hint += "_"
		} else {
			if string(word[wordLength]) == " " {
				hint += "  "
			} else {
				hint += "_ "
			}
		}
	}
	return hint
}
