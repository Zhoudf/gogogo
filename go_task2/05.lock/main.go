package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1
// Counter 包含一个计数器和一个互斥锁
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment 对计数器进行安全的递增操作
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 返回计数器的当前值
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// 1
	var counter Counter
	var wg sync.WaitGroup

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程对计数器进行 1000 次递增操作
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出计数器的最终值
	fmt.Println("Lock计数器的最终值是:", counter.Value())

	// 2
	var counter2 int64
	var wg2 sync.WaitGroup

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			// 每个协程对计数器进行 1000 次递增操作
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter2, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg2.Wait()

	// 输出计数器的最终值
	fmt.Println("atomic计数器的最终值是:", atomic.LoadInt64(&counter2))
}
