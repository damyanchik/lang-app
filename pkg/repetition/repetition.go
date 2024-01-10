package repetition

import (
	"fmt"
	"lang-app/pkg/vocabulary"
)

const (
	KeyOnly   = "KeyOnly"
	ValueOnly = "ValueOnly"
	Mix       = "Mix"
)

type Repetition struct {
	vocabularies  vocabulary.Vocabulary
	countedPoints int
	countedRounds int
	displayMode   string
}

func NewRepetition(vocabularies vocabulary.Vocabulary) Repetition {
	return Repetition{
		vocabularies:  vocabularies,
		countedPoints: 0,
		countedRounds: 0,
		displayMode:   "none",
	}
}

func (r *Repetition) SetDisplayMode(displayMode string) {
	r.displayMode = displayMode
}

func (r *Repetition) ShowPoints() int {
	fmt.Printf("\nTWOJE PUNKTY: %v na %v \n\n", r.countedPoints, len(r.vocabularies.GetVocabularies())*r.countedRounds)
	return r.countedPoints
}

func (r *Repetition) StartRound() {
	isKey := true
	r.countRounds()
	for _, key := range r.vocabularies.GetRandomVocabulariesOrder() {
		value := r.vocabularies.GetVocabularies()[key]

		switch r.displayMode {
		case KeyOnly:
			r.inputString(key, value)
		case ValueOnly:
			r.inputString(value, key)
		case Mix:
			if isKey {
				r.inputString(key, value)
			} else {
				r.inputString(value, key)
			}
			isKey = !isKey
		}
	}
}

func (r *Repetition) inputString(translateWord string, neededWord string) {
	var inputedWord string

	fmt.Printf("Przetłumacz '%v': ", translateWord)
	fmt.Scan(&inputedWord)

	r.countPoints(r.compareWords(inputedWord, neededWord))
}

func (r *Repetition) compareWords(inputedWord string, neededWord string) (bool, string) {
	if inputedWord == neededWord {
		return true, neededWord
	} else {
		return false, neededWord
	}
}

func (r *Repetition) countPoints(point bool, word string) {
	if point {
		r.countedPoints++
		fmt.Println("DOBRZE!")
	} else {
		fmt.Printf("BŁĄD! Poprawne słowo to: %v \n", word)
	}
}

func (r *Repetition) countRounds() {
	r.countedRounds++
	fmt.Printf("## Runda %v! ##\n", r.countedRounds)
}
