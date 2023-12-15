package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

//функция, создающая неотсортированный массив с случайными числами

func arrayCreate() [n]int {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	numbers := [n]int{}
	for i := 0; i < n; i++ {
		numbers[i] = random.Intn(10)
	}
	return numbers
}

func main() {
	currentArray := arrayCreate()

	//анонимная функция, сортирующая массив пузырьком по убыванию

	sortFunc := func(arr [n]int) [n]int {
		for i := 0; i < n-1; i++ {
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		return arr
	}

	fmt.Println("Ваш массив отсортированный по убыванию:", sortFunc(currentArray))
}
