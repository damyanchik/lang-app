package main

import (
	"fmt"
	"lang-app/app/language"
	"lang-app/app/user"
	"lang-app/app/vocabulary"
)

var userName string

func main() {

	language.ChooseLanguage()
	user.GreetUser()
	user.ShowInstruction()

	vocabularies := vocabulary.Process()

	if len(vocabularies) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
	} else {
		fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
		i := 1
		for key, value := range vocabularies {
			fmt.Printf("%v: %v - %v\n", i, key, value)
			i++
		}
	}

}
