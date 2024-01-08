package vocabulary

import (
	"fmt"
	"lang-app/app/helpers"
)

const Ready = "/ready"

var vocabularies = make(map[string]string)

func Process() map[string]string {
	for {
		key := enterWord()

		if key == Ready {
			break
		}

		value := enterTranslation()

		if value == Ready {
			break
		}

		vocabularies[key] = value

		fmt.Printf("#Słowo '%v' i tłumecznie '%v' zostało dodane! Dodaj kolejne lub wpisz /ready\n", key, value)
	}

	return vocabularies
}

func enterWord() string {
	var key string

	for {
		println("\n#Podaj słowo do nauczenia:")
		fmt.Scan(&key)

		if helpers.SearchKeyInMap(key, vocabularies) {
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
