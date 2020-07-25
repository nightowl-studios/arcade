package wordhint

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func countUnderscore(hint string) int {
	count := 0
	for _, char := range hint {
		if string(char) == "_" {
			count++
		}
	}
	return count
}

func TestSpaceInWords(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	hint := wordHint.GiveHint("test s p a c e")
	assertEqual(t, countUnderscore(hint), 7)
}

func TestSpaceInWords2(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	hint := wordHint.GiveHint("test spa c e")
	assertEqual(t, countUnderscore(hint), 7)
}

func TestSpaceInWords3(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	hint := wordHint.GiveHint("test space")
	assertEqual(t, countUnderscore(hint), 7)
}

func TestSpaceInWords4(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	hint := wordHint.GiveHint("t e     s t      s     p     a        c e")
	assertEqual(t, countUnderscore(hint), 7)
}

func TestNumberOfUnderscore(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	hint := wordHint.GiveHint("testing")
	assertEqual(t, countUnderscore(hint), 5)
}

func TestNumberOfUnderscore2(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 4)
	hint := wordHint.GiveHint("testing")
	assertEqual(t, countUnderscore(hint), 4)
}

func TestNumberOfUnderscore3(t *testing.T) {
	wordHint := WordHint{}
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 2)
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 4)
	wordHint.previousHintIndexes = append(wordHint.previousHintIndexes, 6)
	hint := wordHint.GiveHint("testing")
	assertEqual(t, countUnderscore(hint), 3)
}
