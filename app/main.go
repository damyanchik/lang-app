package main

import (
	"fmt"
	"lang-app/app/language"
	"lang-app/app/repetition"
	"lang-app/app/user"
	"lang-app/app/vocabulary"
)

var AppLanguage string

func main() {
	AppLanguage = language.ChooseLanguage()
	user.GreetUser()
	user.ShowInstruction()

	newVocabulary := vocabulary.CreateVocabulary()
	newVocabulary.GetVocabularies()

	if len(newVocabulary.GetVocabularies()) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
		return
	}

	//first round
	newRepetition := repetition.CreateRepetition(newVocabulary, repetition.KeyOnly)
	newRepetition.StartRound()

	fmt.Printf("\nTWOJE PUNKY: %v \n", newRepetition.GetCountedPoints())

	//second round
	newRepetition.SetDisplayMode(repetition.ValueOnly)
	newRepetition.StartRound()
	fmt.Printf("\nTWOJE PUNKY: %v \n", newRepetition.GetCountedPoints())

	//third round
	newRepetition.SetDisplayMode(repetition.Mix)
	newRepetition.StartRound()

	fmt.Println("KONIEC")
	fmt.Printf("\nZDOBYŁEŚ ŁĄCZNIE: %v \n", newRepetition.GetCountedPoints())
}
