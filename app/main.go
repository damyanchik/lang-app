package main

import (
	"fmt"
	"lang-app/app/language"
	"lang-app/app/repetition"
	"lang-app/app/user"
	"lang-app/app/vocabulary"
)

func main() {
	language.NewAppLanguage()
	user.GreetUser()
	user.ShowInstruction()

	newVocabulary := vocabulary.NewVocabulary()
	newVocabulary.GetVocabularies()

	if len(newVocabulary.GetVocabularies()) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
		return
	}

	newRepetition := repetition.NewRepetition(newVocabulary)
	for _, roundType := range [3]string{repetition.KeyOnly, repetition.ValueOnly, repetition.Mix} {
		newRepetition.SetDisplayMode(roundType)
		newRepetition.StartRound()
		newRepetition.ShowPoints()
	}

	fmt.Println("KONIEC!")
}
