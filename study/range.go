package study

import "fmt"

func RangeTest() {
	arr := []int{12, 33, 444, 5}

	sum := 0
	for _, num := range arr {
		sum += num
	}

	fmt.Println(sum)

	fruit := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range fruit {
		fmt.Printf("%s -> %s \n", k, v)
	}
}
