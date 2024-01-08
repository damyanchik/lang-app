package language

import (
	"fmt"
	"lang-app/app/helpers"
	"strings"
)

var appLanguage string

func ChooseLanguage() {
	availableLanguages := []string{"EN", "PL", "RU"}
	fmt.Println("#Wybierz język aplikacji: (EN / PL / RU)")

	for {
		var chosenLanguage string

		fmt.Scan(&chosenLanguage)
		appLanguage = strings.ToUpper(chosenLanguage)

		if helpers.IsInArray(appLanguage, availableLanguages) {
			break
		} else {
			fmt.Println("#Język został wybrany niepoprawnie, spróbuj ponownie!")
		}
	}
}
