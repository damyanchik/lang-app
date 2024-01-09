package vocabulary

import (
	"fmt"
	"lang-app/app/collections"
)

const Ready = "/ready"

var vocabularies = make(map[string]string)

func GetRandomVocabulariesOrder(vocabularies map[string]string) []string {
	keyCollection := collections.MapKeysToSlice(vocabularies)
	newCollectionOrder := collections.RandomOrderInSlice(keyCollection)

	return newCollectionOrder
}

func Preparation() map[string]string {
	for len(vocabularies) < 20 {
		key := enterWord()

		if key == Ready {
			break
		}

		value := enterTranslation()

		if value == Ready {
			break
		}

		vocabularies[key] = value

		fmt.Printf("\n#Słowo '%v' i tłumecznie '%v' zostało dodane! Możesz dodać jeszcze %v słówek.\n", key, value, 20-len(vocabularies))
		fmt.Println("Dodaj kolejne słówko lub wpisz /ready, aby zakończyć")
	}

	return vocabularies
}

func enterWord() string {
	var key string

	for {
		println("\n#Podaj słowo do nauczenia:")
		fmt.Scan(&key)

		if collections.IsKeyInMap(key, vocabularies) {
			println("!!! Podane słówko było już podane wcześniej! Podaj inne. !!!")
			continue
		}

		break
	}

	return key
}

func enterTranslation() string {
	var value string

	fmt.Println("#Podaj tłumaczenie słowa:")
	fmt.Scan(&value)

	return value
}
