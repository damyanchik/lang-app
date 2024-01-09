package vocabulary

import (
	"fmt"
	"lang-app/app/collections"
)

const Ready = "/ready"

type Vocabulary struct {
	vocabularies map[string]string
	wordLimit    int
}

func CreateVocabulary(wordLimit int) Vocabulary {
	v := Vocabulary{
		wordLimit: wordLimit,
	}
	v.preparationVocabularies()

	return v
}

func (v *Vocabulary) GetVocabularies() map[string]string {
	return v.vocabularies
}

func (v *Vocabulary) GetRandomVocabulariesOrder() []string {
	keysFromVocabularies := collections.MapKeysToSlice(v.vocabularies)
	newVocabulariesOrder := collections.RandomOrderInSlice(keysFromVocabularies)

	return newVocabulariesOrder
}

func (v *Vocabulary) preparationVocabularies() {
	for len(v.vocabularies) < v.wordLimit {
		key := v.enterWord()

		if v.isReady(key) {
			break
		}

		value := v.enterTranslation()

		if v.isReady(value) {
			break
		}

		v.addVocabulary(key, value)
	}
	v.vocabularyInformation()
}

func (v *Vocabulary) enterWord() string {
	var key string

	for {
		println("\n#Podaj słowo do nauczenia:")
		fmt.Scan(&key)

		if collections.IsKeyInMap(key, v.vocabularies) {
			println("!!! Podane słówko było już podane wcześniej! Podaj inne. !!!")
			continue
		}

		break
	}

	return key
}

func (v *Vocabulary) enterTranslation() string {
	var value string

	fmt.Println("#Podaj tłumaczenie słowa:")
	fmt.Scan(&value)

	return value
}

func (v *Vocabulary) addVocabulary(key string, value string) {
	v.vocabularies[key] = value
	fmt.Printf("\n#Słowo '%v' i tłumecznie '%v' zostało dodane! Możesz dodać jeszcze %v słówek.\n", key, value, 20-len(v.vocabularies))
	fmt.Println("Dodaj kolejne słówko lub wpisz /ready, aby zakończyć")
}

func (v *Vocabulary) vocabularyInformation() {
	fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
	collections.PrintMap(v.GetVocabularies())
}

func (v *Vocabulary) isReady(input string) bool {
	if input != Ready {
		return false
	} else {
		return true
	}
}
