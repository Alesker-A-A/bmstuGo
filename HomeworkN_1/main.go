package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 5)
	sum := 0

	fmt.Print("Введите 5 натуральных чисел: ")
	for i := range nums {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	fmt.Println("Отсортированные элементы: ", nums)
	fmt.Println("Самое большое число: ", nums[0])
	fmt.Println("Самое маленькое число: ", nums[4])
	fmt.Printf("Среднее арифметическое: %.2f\n", float64(sum)/5)
}
