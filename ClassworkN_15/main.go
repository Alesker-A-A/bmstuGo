// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	files, err := os.ReadDir(".")
// 	if err != nil {
// 		fmt.Println("Cannot readdir =(\n", err)
// 		return
// 	}
// 	for _, file := range files {
// 		fmt.Println(file.Name())
// 	}
// }

// ОБРАБОТКА ОШИБОК
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// type NewWriter struct {
// 	w   io.Writer
// 	err error
// }

// func (nw *NewWriter) WriteNewLine(line string) {
// 	if nw.err != nil {
// 		return
// 	}
// 	_, nw.err = fmt.Fprintln(nw.w, line)
// }
// func WriteToFile(name string) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	defer f.Close()
// 	nw := NewWriter{
// 		w: f,
// 	}
// 	nw.WriteNewLine("lesson 15 go")
// 	nw.WriteNewLine("hello")
// 	nw.WriteNewLine("This")
// 	nw.WriteNewLine("is")
// 	nw.WriteNewLine("best")
// 	nw.WriteNewLine("BMSTU")
// 	nw.WriteNewLine("Cource")
// 	nw.WriteNewLine("that's all")
// 	return nw.err
// }

// func main() {
// 	err := WriteToFile("test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("All is OK!")
// }

// СОЗДАНИЕ СВОИХ ОШИБОК
// package main

// import (
// 	"fmt"
// )

// type newError struct{}

// func (newerr newError) Error() string {
// 	return "This is new Error"
// }

// func someFunc() error {
// 	var NextGenError newError
// 	return NextGenError
// }

// func main() {
// 	err := fmt.Errorf("some error")
// 	// err = errors.New(fmt.Sprintf("Another error"))

// }

// ОБЕРТКИ ОШИБОК/ПОСЛЕДОВАТЕЛЬНЫЕ ОШИБКИ
package main

import "fmt"

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("Error while calling f2: %w", err)
	}
	return nil
}
func f2() error {
	err := f3() //
	if err != nil {
		return fmt.Errorf("Error while calling f3: %w", err)
	}
	return nil
}
func f3() error {
	return fmt.Errorf("I'm error in f3")
}
func main() {
	err := f1()
	if err != nil {
		fmt.Println(err)
	}
}
