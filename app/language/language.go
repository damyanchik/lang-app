package language

import (
	"fmt"
	"lang-app/app/collections"
	"strings"
)

const (
	EN = "EN"
	PL = "PL"
	RU = "RU"
)

var appLanguage string
var availableLanguages = []string{EN, PL, RU}

func ChooseLanguage() string {
	fmt.Printf("#Wybierz język aplikacji: %v \n", availableLanguages)

	for {
		var chosenLanguage string

		fmt.Scan(&chosenLanguage)
		appLanguage = strings.ToUpper(chosenLanguage)

		if collections.IsInArray(appLanguage, availableLanguages) {
			break
		} else {
			fmt.Println("#Język został wybrany niepoprawnie, spróbuj ponownie!")
		}
	}

	return appLanguage
}
