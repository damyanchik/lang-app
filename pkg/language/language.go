package language

import (
	"fmt"
	"lang-app/pkg/collections"
	"strings"
)

const (
	//EN = "EN"
	PL = "PL"
	//RU = "RU"
)

type AppLanguage struct {
	availableLanguages []string
	chosenLanguage     string
}

func NewAppLanguage() AppLanguage {
	a := AppLanguage{
		availableLanguages: []string{PL},
	}

	fmt.Printf("#Wybierz język aplikacji: %v \n", a.availableLanguages)
	a.chooseLanguage()

	return a
}

func (a *AppLanguage) GetChosenLanguage() string {
	return a.chosenLanguage
}

func (a *AppLanguage) chooseLanguage() {
	for {
		var inputedLanguage string

		fmt.Scan(&inputedLanguage)
		appLanguage := strings.ToUpper(inputedLanguage)

		if collections.IsInArray(appLanguage, a.availableLanguages) {
			a.chosenLanguage = appLanguage
			break
		} else {
			fmt.Println("#Język został wybrany niepoprawnie, spróbuj ponownie!")
		}
	}
}
