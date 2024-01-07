package main

import (
	"fmt"
	"strings"
)

var appLanguage string

func main() {
	var vocabulary = make(map[string]string)

	fmt.Println("#Witaj w programie do nauki języka LangApp!")
	fmt.Println("#Wybierz język aplikacji: (EN / PL / RU)")

	for {
		var chosenLanguage string
		availableLanguages := []string{"EN", "PL", "RU"}

		fmt.Scan(&chosenLanguage)
		appLanguage = strings.ToUpper(chosenLanguage)

		if isInArray(appLanguage, availableLanguages) {
			break
		} else {
			fmt.Println("#Język został wybrany niepoprawnie, spróbuj ponownie!")
		}
	}

	fmt.Println("#Ekspeci sugerują, aby skupić się każdego dnia na nauce od 10 do 20 nowych słówek. Aplikacja pozwoli Ci na naukę maksymalnie 20 słówek.")
	fmt.Println("#Najpierw zapisz słowo w języku obcym, którego się uczysz, a następnie zostaniesz poproszony o podanie jego tłumaczenia.")
	fmt.Println("#Możesz podać maksymalnie 20 słówek")

	for {
		var key, value string

		println("#Podaj słowo do nauczenia:")
		fmt.Scan(&key)

		if key == "/ready" {
			break
		}

		fmt.Printf("#Podaj tłumaczenie słowa '%v':\n", key)
		fmt.Scan(&value)

		if value == "/ready" {
			break
		}

		vocabulary[key] = value

		fmt.Printf("#Słowo '%v' i tłumecznie '%v' zostało dodane! Dodaj kolejne lub wpisz /ready\n", key, value)
	}

	if len(vocabulary) == 0 {
		fmt.Println("#Brak słówek do nauczenia, koniec działania programu.")
	} else {
		fmt.Println("#Zakończono dodawanie słówek, oto twoja lista: (SŁOWO - TŁUMACZENIE)")
		i := 1
		for key, value := range vocabulary {
			fmt.Printf("%v: %v - %v\n", i, key, value)
			i++
		}
	}

}

func isInArray(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

//ukaż ile zostało już wpisanych
//usuń
//edytuj
//pokaż listę słówek
//słowa nie mogą się powtarzać
//EN PL RU FR
//jezeli

//Instrukcja i Dodaj fiszki

//
