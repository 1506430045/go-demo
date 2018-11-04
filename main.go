package main

import "fmt"

func calc(ch chan int) {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum++
	}
	ch <- sum
}

func main() {

	sum1ch := make(chan int)
	sum2ch := make(chan int)

	go calc(sum1ch)
	go calc(sum2ch)

	sum1, sum2 := <-sum1ch, <-sum2ch
	fmt.Println(sum1 + sum2)

	//study.ArrayTest()
	//study.PointerTest()
	//book := study.Book{
	//	Title:   "钢铁是怎样炼成的",
	//	Author:  "鲁迅",
	//	Subject: "文学",
	//	Id:      1,
	//}
	//study.PrintBook(book)
	//study.SliceTest()
	//study.RangeTest()
	//study.MapTest()
	//study.RecursionTest()
	//study.InterfaceTest()
	//study.GoroutineTest()
	//study.ChannelTest()
	//study.Channel4Test()

	//defer db.SqlDB.Close()
	//router := initRouter()
	//router.Run(":8018")
}
