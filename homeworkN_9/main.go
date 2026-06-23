package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль невозможно!")
	}
	return a / b, nil
}

func main() {
	var a float64
	var b float64

	fmt.Printf("Введите значение для а: ")
	fmt.Scan(&a)

	fmt.Printf("Введите значение для b: ")
	fmt.Scan(&b)

	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
