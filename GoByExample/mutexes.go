// 在前面的例子中，我们看到了如何使用原子操作(atomic-counters)来管理简单的计数器。
// 对于更加复杂的情况，我们可以使用一个互斥量 来在 Go 协程间安全的访问数据。
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Container{
		// 请注意，互斥量的零值是可用的，因此这里不需要初始化。
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// 这个函数在循环中递增对 name 的计数
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)

	// 同时运行多个 goroutines; 请注意，它们都访问相同的 Container，其中两个访问相同的计数器。
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// 等待上面的 goroutines 都执行结束
	wg.Wait()
	fmt.Println(c.counters)
}

// Container 中定义了 counters 的 map ，由于我们希望从多个 goroutine 同时更新它，
// 因此我们添加了一个 互斥锁Mutex 来同步访问。
// 请注意不能复制互斥锁，如果需要传递这个 struct，应使用指针完成。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()

	// 在访问 counters 之前锁定互斥锁； 使用 [defer]（defer） 在函数结束时解锁。
	defer c.mu.Unlock()
	c.counters[name]++
}

// go run mutexes.go
// 输出：
// map[a:20000 b:10000]
