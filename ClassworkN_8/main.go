package main

import "fmt"

type Animal struct {
	Name  string
	Voice string
	aType string
}

func (a Animal) talk() {
	fmt.Printf("%s %s говорит: %s\n", a.aType, a.Name, a.Voice)
}

func main() {
	Barsic := Animal{
		Name:  "Барсик",
		Voice: "Мяу!",
		aType: "Кот",
	}
	Sharic := Animal{
		Name:  "Шарик",
		Voice: "Гав!",
		aType: "Пес",
	}
	Barsic.talk()
	Sharic.talk()
}
