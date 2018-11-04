package study

import "fmt"

func ArrayTest() {
	arr := []int{1, 2, 3, 4, 5}

	avg := getAvg(arr, 5)

	fmt.Println(avg)
}

func getAvg(arr []int, size int) float32 {
	var sum = 0
	var avg float32 = 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
