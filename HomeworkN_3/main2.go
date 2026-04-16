package main

import "fmt"

type City struct {
	cityName string
}

var city []City

func showMenu() {
	fmt.Println("1. Показать города")
	fmt.Println("2. Добавить город")
	fmt.Println("3. Удалить город")
	fmt.Println("4. Выход")
}

func showCities() {
	if len(city) == 0 {
		fmt.Println("Список городов пуст")
		return
	}
	fmt.Println("*** Список городов ***")
	for _, c := range city {
		fmt.Println(c.cityName)
	}
}

func addCity() {
	var name string
	fmt.Print("Введите название города: ")
	fmt.Scan(&name)

	city = append(city, City{cityName: name})
}

func deleteCity() {
	var name string
	fmt.Print("Введите название города для удаления: ")
	fmt.Scan(&name)

	for i, c := range city {
		if c.cityName == name {
			city = append(city[:i], city[i+1:]...)
			fmt.Println("Город удален!")
			return
		}
	}
	fmt.Println("Город не найден :(")
}

func main() {
	for {
		showMenu()
		var choice int
		fmt.Scan(&choice)

		// использую тут switch case так как визуально чище и проще читать код
		//ну и лень было расписывать if/else if

		switch choice {
		case 1:
			showCities()
		case 2:
			addCity()
		case 3:
			deleteCity()
		case 4:
			fmt.Println("До свидания!")
			return
		}
	}
}
