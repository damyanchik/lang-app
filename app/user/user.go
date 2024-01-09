package user

import (
	"fmt"
	"lang-app/app/vocabulary"
	"time"
)

var userName string

func GreetUser() {
	fmt.Println("#Witaj w programie do nauki języka LangApp!")
	fmt.Println("#Podaj nam swoje imię:")
	fmt.Scan(&userName)
	fmt.Printf("\n##Witaj %v!\n", userName)
	fmt.Printf("#Ekspeci sugerują, aby skupić się każdego dnia na nauce od 10 do 20 nowych słówek. Aplikacja pozwoli Ci na naukę maksymalnie %v słówek. Poniżej przedstawiamy instrukcję. \n", vocabulary.WordLimit)
}

func ShowInstruction() {
	time.Sleep(2 * time.Second)
	fmt.Println("\n## INSTRUKCJA ##")
	fmt.Println("#Najpierw zapisz słowo w języku obcym, którego się uczysz, a następnie podaj jego tłumaczenie.")
	fmt.Printf("#Możesz podać maksymalnie %v słówek. \n", vocabulary.WordLimit)
}
