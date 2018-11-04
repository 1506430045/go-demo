package study

//使用信道是要考虑的一个重要因素是死锁（Deadlock）。如果一个协程发送数据给一个信道，而没有其他的协程从该信道中接收数据，那么程序在运行时会遇到死锁，并触发 panic 。
func Channel3Test() {
	ch := make(chan int)
	ch <- 5
}
