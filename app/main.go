package main

import (
	"fmt"
	"lang-app/app/collections"
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

	vocabularies := vocabulary.Preparation()

	if len(vocabularies) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
		return
	}

	vocabularyInformation(vocabularies)

	//first round
	fmt.Println("## Runda pierwsza ##")
	newRepetition := repetition.CreateRepetition(vocabularies, repetition.KeyOnly)
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

func vocabularyInformation(vocabularies map[string]string) {
	fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
	collections.PrintMap(vocabularies)
	fmt.Println("#Teraz poczekaj chwilę, zaraz rozpoczniemy naukę! Przed nami trzy rundy.")
}
