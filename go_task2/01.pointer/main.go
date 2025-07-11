package main

import "fmt"

// 1. incrementByTen 接收一个整数指针作为参数，将该指针指向的值增加 10
func incrementByTen(numPtr *int) {
	// 通过指针解引用获取指向的值，并增加 10
	*numPtr += 10
}

// 2. multiplySliceByTwo 接收一个整数切片的指针，将切片中的每个元素乘以 2
func multiplySliceByTwo(slicePtr *[]int) {
	// 通过指针解引用获取切片
	slice := *slicePtr
	for i := range slice {
		// 将切片中的每个元素乘以 2
		slice[i] *= 2
	}
}

func main() {
	// 1
	// 定义一个整数变量
	num := 5
	fmt.Printf("增加前的值: %d\n", num)

	// 将变量的地址传递给 incrementByTen 函数
	incrementByTen(&num)

	// 输出增加后的值
	fmt.Printf("增加后的值: %d\n", num)

	// 2
	// 定义一个整数切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("乘以 2 前的切片: %v\n", slice)

	// 将切片的地址传递给 multiplySliceByTwo 函数
	multiplySliceByTwo(&slice)

	// 输出乘以 2 后的切片
	fmt.Printf("乘以 2 后的切片: %v\n", slice)
}
