package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	vowels := "аиеёийоуыэюя"
	count := 0

	for _, ch := range input {
		if strings.ContainsRune(vowels, ch) {
			count = count + 1
		}
	}

	fmt.Println("Количество гласных: ", count)
}
