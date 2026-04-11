package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkBrackets(s string) (bool, int, int) {
	open := 0
	close := 0
	balance := 0

	for _, ch := range s {
		if ch == '(' {
			open++
			balance++
		} else if ch == ')' {
			close++
			balance--
		}
		if balance < 0 {
			return false, open, close
		}
	}
	return balance == 0, open, close
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите формулу: ")
	scanner.Scan()
	input := scanner.Text()

	valid, open, close := checkBrackets(input)

	if valid {
		fmt.Printf("Скобки расставлены верно, %d открывающихся, %d закрывющихся\n", open, close)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающихся, %d закрывающихся\n", open, close)
	}
}
