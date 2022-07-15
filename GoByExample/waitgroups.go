// 想要等待多个协程完成，我们可以使用 **wait group**。
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 这个 WaitGroup 用于等待这里启动的所有协程完成。
	// 注意：如果 WaitGroup 显式传递到函数中，则应使用 *指针*。
	var wg sync.WaitGroup

	// 启动几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 也可以在 for 外层 wg.add(5)

		// 避免在每个协程闭包中重复利用相同的 i 值 更多细节可以参考 the FAQ
		i := 1

		// 将 worker 调用包装在一个闭包中，可以确保通知 WaitGroup 此工作线程已完成。
		// 这样，worker 线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			defer wg.Done()
			newWorker(i)
		}()
	}

	// 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成。
	wg.Wait()
}

// 每个协程都会运行该函数。
func newWorker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// 睡眠一秒钟，以此来模拟耗时的任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
