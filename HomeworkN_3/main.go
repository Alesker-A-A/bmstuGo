package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arr [10]int
	for i := range arr {
		arr[i] = rand.Intn(100) + 1
	}

	slice := make([]int, len(arr))
	copy(slice, arr[:])
	sort.Ints(slice)
	fmt.Println("Исходный массив: ", arr)
	fmt.Println("Отсортированный срез: ", slice)
}
