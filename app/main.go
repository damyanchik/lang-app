package main

import (
	"fmt"
	"lang-app/app/collections"
	"lang-app/app/language"
	"lang-app/app/user"
	"lang-app/app/vocabulary"
)

func main() {

	language.ChooseLanguage()
	user.GreetUser()
	user.ShowInstruction()

	vocabularies := vocabulary.Process()

	if len(vocabularies) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
		return
	} else {
		fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
		collections.PrintMap(vocabularies)
	}
}
