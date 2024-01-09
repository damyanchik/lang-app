package main

import (
	"fmt"
	"lang-app/app/collections"
	"lang-app/app/game"
	"lang-app/app/language"
	"lang-app/app/user"
	"lang-app/app/vocabulary"
)

func main() {

	language.ChooseLanguage()
	user.GreetUser()
	user.ShowInstruction()

	vocabularies := vocabulary.Preparation()

	if len(vocabularies) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
		return
	}

	fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
	collections.PrintMap(vocabularies)
	fmt.Println("#Teraz poczekaj chwilę, zaraz rozpoczniemy naukę! Przed nami trzy rundy.")

	fmt.Println("## Runda pierwsza ##")

	points := game.Play(vocabularies, game.KeyOnly)

	fmt.Printf("\nTWOJE PUNKY: %v \n", points)
	fmt.Println("## Runda druga ##")

	points += game.Play(vocabularies, game.ValueOnly)

	fmt.Printf("\nTWOJE PUNKY: %v \n", points)
	fmt.Println("## Runda trzecia ##")

	points += game.Play(vocabularies, game.Mix)

	fmt.Println("KONIEC")
	fmt.Printf("\nZDOBYŁEŚ ŁĄCZNIE: %v \n", points)
}
