package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalizeWords(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		runes := []rune(word)
		if len(runes) == 0 {
			continue
		}
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Println(capitalizeWords(input))
}
