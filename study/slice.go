package study

import "fmt"

func SliceTest() {

	s1 := make([]int, 3, 5)
	s1 = []int{21, 33, 232, 22, 2, 11}

	s1 = append(s1, 44)

	s2 := make([]int, len(s1), cap(s1)*2)

	copy(s2, s1)

	fmt.Println(len(s1), cap(s1), s1, s2)
}
