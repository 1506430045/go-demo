package study

import "fmt"

func RecursionTest() {
	fmt.Println(factorial(4))

	var i int64
	for i = 1; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

func factorial(n uint64) uint64 {
	if n > 1 {
		return n * factorial(n-1)
	}
	return 1
}

func fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
