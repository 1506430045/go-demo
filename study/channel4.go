package study

import "fmt"

func Channel4Test() {
	ch := make(chan int)

	//go producer(ch)
	//
	//for v := range ch {
	//	fmt.Println("Received ", v)
	//}
	//
	//fmt.Println("---------------")

	go producer(ch)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
