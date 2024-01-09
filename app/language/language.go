package language

import (
	"fmt"
	"lang-app/app/collections"
	"strings"
)

var appLanguage string

func ChooseLanguage() string {
	availableLanguages := []string{"EN", "PL", "RU"}
	fmt.Println("#Wybierz język aplikacji: (EN / PL / RU)")

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
