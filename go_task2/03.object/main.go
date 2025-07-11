package main

import (
	"fmt"
	"math"
)

// 1
// Shape 定义 Shape 接口，包含 Area() 和 Perimeter() 两个方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 定义矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形的面积，实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形的周长，实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 定义圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形的面积，实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形的周长，实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 2
// Person 定义 Person 结构体，包含 Name 和 Age 字段
type Person struct {
	Name string
	Age  int
}

// Employee 定义 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段
type Employee struct {
	Person     //匿名时可以直接调用。起名时需要使用类似e.person.Name
	EmployeeID string
}

// PrintInfo 为 Employee 结构体实现 PrintInfo 方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工 ID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	// 1
	// 创建 Rectangle 结构体实例
	rect := Rectangle{Width: 5, Height: 10}
	// 创建 Circle 结构体实例
	circle := Circle{Radius: 3}

	// 调用 Rectangle 的 Area 和 Perimeter 方法
	fmt.Printf("矩形的面积: %.2f\n", rect.Area())
	fmt.Printf("矩形的周长: %.2f\n", rect.Perimeter())

	// 调用 Circle 的 Area 和 Perimeter 方法
	fmt.Printf("圆形的面积: %.2f\n", circle.Area())
	fmt.Printf("圆形的周长: %.2f\n", circle.Perimeter())

	// 2
	// 创建 Employee 结构体实例
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E001",
	}

	// 调用 PrintInfo 方法输出员工信息
	employee.PrintInfo()
}
