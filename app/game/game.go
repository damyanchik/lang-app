package game

import (
	"fmt"
	"lang-app/app/collections"
)

const (
	KeyOnly   = "key"
	ValueOnly = "value"
	Mix       = "mix"
)

func Play(vocabularies map[string]string, displayMode string) int {
	var points int
	isKey := true

	keyCollection := collections.MapKeysToSlice(vocabularies)
	newCollectionOrder := collections.RandomOrderInSlice(keyCollection)

	for _, key := range newCollectionOrder {
		value := vocabularies[key]

		switch displayMode {
		case KeyOnly:
			points += inputString(key, value)
		case ValueOnly:
			points += inputString(value, key)
		case Mix:
			if isKey {
				points += inputString(key, value)
			} else {
				points += inputString(value, key)
			}
			isKey = !isKey
		}
	}

	return points
}

func inputString(translateWord string, guessWord string) int {
	var str string

	fmt.Printf("Przetłumacz '%v': ", translateWord)
	fmt.Scan(&str)

	return verifyWordInMap(str, guessWord)
}

func verifyWordInMap(str string, word string) int {
	point := 0
	if str == word {
		point += 1
		fmt.Println("DOBRZE!")
	} else {
		fmt.Printf("BŁĄD! Poprawne słowo to: %v \n", word)
	}

	return point
}
