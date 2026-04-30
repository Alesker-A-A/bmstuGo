// package main

// import "fmt"

// func kettle(ch chan string) {
// 	fmt.Println("Чайник запущен")
// 	ch <- "чайник готов" // отправляем сигнал
// }

// func coffeeMachine(ch chan string) {
// 	message := <-ch // ждем сигнал от чайника
// 	fmt.Println("Получен сигнал: ", message)
// 	fmt.Println("Кофеварка запущена!")
// }

// func main() {
// 	ch := make(chan string)
// 	go kettle(ch)
// 	go coffeeMachine(ch)

//		var input string
//		fmt.Scan(&input)
//	}
// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func generate(ch chan int) {
// 	number := rand.Intn(100) + 1
// 	ch <- number
// }

//	func main() {
//		ch := make(chan int)
//		go generate(ch)
//		result := <-ch
//		fmt.Println("Получено:", result)
//	}
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Горутина: ", n)
		}(i)
		wg.Wait()
	}
}
