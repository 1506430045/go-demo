package study

import "fmt"

func ChannelTest() {
	a := []int{1, 2, 4, 7, 8, 2}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, num := range a {
		sum += num
	}
	c <- sum
}
