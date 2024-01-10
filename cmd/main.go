package main

import (
	"fmt"
	"lang-app/pkg/language"
	"lang-app/pkg/repetition"
	"lang-app/pkg/user"
	"lang-app/pkg/vocabulary"
	"time"
)

func main() {
	language.NewAppLanguage()
	user.NewUser()
	ShowInstruction()

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

func ShowInstruction() {
	time.Sleep(2 * time.Second)
	fmt.Println("\n## INSTRUKCJA ##")
	fmt.Println("#Najpierw zapisz słowo w języku obcym, którego się uczysz, a następnie podaj jego tłumaczenie.")
	fmt.Printf("#Możesz podać maksymalnie %v słówek. \n", vocabulary.WordLimit)
}
