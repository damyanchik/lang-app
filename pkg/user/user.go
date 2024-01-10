package user

import (
	"fmt"
	"lang-app/pkg/vocabulary"
)

type User struct {
	name string
}

func NewUser() User {
	u := User{}
	u.greetUser()

	return u
}

func (u *User) greetUser() {
	var userName string
	fmt.Println("#Witaj w programie do nauki języka LangApp!")
	fmt.Println("#Podaj nam swoje imię:")
	fmt.Scan(&userName)
	u.name = userName
	fmt.Printf("\n##Witaj %v!\n", userName)
	fmt.Printf("#Ekspeci sugerują, aby skupić się każdego dnia na nauce od 10 do 20 nowych słówek. Aplikacja pozwoli Ci na naukę maksymalnie %v słówek. \n", vocabulary.WordLimit)
}
