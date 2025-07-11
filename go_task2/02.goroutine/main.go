package main

import (
	"fmt"
	"sync"
	"time"
)

// 1
// printOdd 函数用于打印 1 到 10 的奇数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知 WaitGroup 任务完成
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// printEven 函数用于打印 2 到 10 的偶数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知 WaitGroup 任务完成
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// 2
// Task 定义任务类型，为无参数、无返回值的函数
type Task func()

// TaskScheduler 任务调度器结构体
type TaskScheduler struct {
	wg sync.WaitGroup
}

// Schedule 方法用于调度一组任务并发执行
func (ts *TaskScheduler) Schedule(tasks []Task) {
	for i, task := range tasks {
		ts.wg.Add(1)
		// 启动协程执行任务
		go ts.runTask(i, task)
	}
	// 等待所有任务完成
	ts.wg.Wait()
}

// runTask 方法用于执行单个任务并统计执行时间
func (ts *TaskScheduler) runTask(taskIndex int, task Task) {
	defer ts.wg.Done()
	start := time.Now()
	// 执行任务
	task()
	duration := time.Since(start)
	fmt.Printf("任务 %d 执行时间: %v\n", taskIndex+1, duration)
}

func main() {
	//1
	var wg sync.WaitGroup
	wg.Add(2) // 告诉 WaitGroup 有两个协程需要等待

	// 启动打印奇数的协程
	go printOdd(&wg)
	// 启动打印偶数的协程
	go printEven(&wg)

	wg.Wait() // 等待所有协程执行完毕

	//2
	// 定义示例任务
	tasks := []Task{
		func() {
			time.Sleep(2 * time.Second)
		},
		func() {
			time.Sleep(1 * time.Second)
		},
		func() {
			time.Sleep(1500 * time.Millisecond)
		},
	}

	// 创建任务调度器实例
	scheduler := TaskScheduler{}
	// 调度任务执行
	scheduler.Schedule(tasks)
}
