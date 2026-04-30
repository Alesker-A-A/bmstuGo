package main

import (
	"fmt"
	"os"
)

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter message: ")
// input, err := reader.ReadString('\n')
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(input)
//fmt.Fprintln(os.Stdout, "hello", "world!")

func main() {
	// 	_, err := os.Create("test.txt")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
