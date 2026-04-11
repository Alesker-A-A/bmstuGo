package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Println("Колличество символов: ", len(input))
}
