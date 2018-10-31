package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 1)
	timeout := make(chan bool, 1)
	// 并发执行一个函数，等待1s后向timeout写入true
	go func() {
		time.Sleep(1000)
		timeout <- true
	}()
	// 这里会等待ch或timeout读出数据
	// 因为一直没有向ch写入数据
	// 在1s后向timeout写入了数据
	// 所以执行了timeout的case
	// 利用这个技巧可以实现超时操作
	select {
	case <-ch:
		fmt.Println("read from ch")
	case <-timeout:
		fmt.Println("time out...")
	}
}
