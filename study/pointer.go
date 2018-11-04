package study

import "fmt"

const MAX = 3

func PointerTest() {

	var a = [MAX]int{1, 4, 7}
	var ptr [MAX] *int

	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
