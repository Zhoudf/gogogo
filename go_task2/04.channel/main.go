package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1
	// 创建一个无缓冲的整数通道
	ch := make(chan int)
	var wg sync.WaitGroup

	// 启动发送协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 生成 1 到 10 的整数并发送到通道
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		// 关闭通道，表示数据发送完毕
		close(ch)
	}()

	// 启动接收协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 从通道接收数据并打印
		for num := range ch {
			fmt.Println(num)
		}
	}()

	// 等待两个协程完成
	wg.Wait()

	//2
	// 创建一个带有缓冲的整数通道，缓冲区大小为 100
	ch2 := make(chan int, 100)
	var wg2 sync.WaitGroup

	// 启动生产者协程
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		// 生成 1 到 100 的整数并发送到通道
		for i := 1; i <= 100; i++ {
			ch2 <- i
		}
		// 关闭通道，表示数据发送完毕
		close(ch2)
	}()

	// 启动消费者协程
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		// 从通道接收数据并打印
		for num := range ch2 {
			fmt.Println(num)
		}
	}()

	// 等待生产者和消费者协程完成
	wg2.Wait()
}
