// 我们可以使用通道来同步协程之间的执行状态。
// 这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。
// 如果需要等待多个协程，WaitGroup 是一个更好的选择。
package main

import (
	"fmt"
	"time"
)

func main() {
	// 运行一个 worker 协程，并给予用于通知的通道。
	d := make(chan bool, 1)
	go worker(d)

	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	<-d
	// 如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束了。
}

// 我们将要在协程中运行这个函数。
// done 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Print("waiting...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true
}

// go run channel_synchronization.go
// 输出：
// waiting...done
