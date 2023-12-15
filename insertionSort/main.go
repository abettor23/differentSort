package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

//функция, создающая и сортирующая массив алгоритмом вставки

func arrayCreate() [n]int {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	numbers := [n]int{}
	for i := 0; i < n; i++ {
		numbers[i] = random.Intn(10)
	}

	for i := 1; i < n; i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = key
	}
	return numbers
}

func main() {
	fmt.Println("Отсортированный массив:", arrayCreate())
}
