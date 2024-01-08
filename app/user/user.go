package user

import (
	"fmt"
	"time"
)

var userName string

func GreetUser() {
	fmt.Println("#Witaj w programie do nauki języka LangApp!")
	fmt.Println("#Podaj nam swoje imię:")
	fmt.Scan(&userName)
	fmt.Printf("##Witaj %v!\n", userName)
	fmt.Println("#Ekspeci sugerują, aby skupić się każdego dnia na nauce od 10 do 20 nowych słówek. Aplikacja pozwoli Ci na naukę maksymalnie 20 słówek. Poniżej przedstawiamy instrukcję.")
}

func ShowInstruction() {
	time.Sleep(2 * time.Second)
	fmt.Println("## INSTRUKCJA ##")
	fmt.Println("#Najpierw zapisz słowo w języku obcym, którego się uczysz, a następnie zostaniesz poproszony o podanie jego tłumaczenia.")
	fmt.Println("#Możesz podać maksymalnie 20 słówek")
}
