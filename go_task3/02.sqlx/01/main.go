package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlx与原生包对比
// 结构体映射：sqlx 支持自动映射，database/sql 需手动扫描。
// 便捷查询方法：sqlx 提供 Get 和 Select 等方法并自动处理 ErrNoRows 错误。
// 命名参数支持：sqlx 支持命名参数，database/sql 只支持占位符。
// 预处理语句缓存：sqlx 内置缓存机制，database/sql 无内置缓存。

// Employee 定义员工结构体，与 employees 表字段对应
// 不同点 1: 结构体映射
// - database/sql: 需要手动逐列扫描查询结果到变量，代码冗长且易出错。
// - sqlx: 支持自动将查询结果映射到结构体，通过 db tag 指定字段对应关系，代码简洁。
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// getDBConnection 获取数据库连接
func getDBConnection() (*sqlx.DB, error) {
	// 数据库连接信息，需要根据实际情况修改
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	// 不同点 5: 预处理语句缓存
	// - database/sql: 没有内置预处理语句缓存机制，需手动管理。
	// - sqlx: 内置预处理语句缓存，使用 Preparex 或 PrepareNamed 创建的预处理语句会被缓存，提升性能。
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// getTechEmployees 查询所有部门为 "技术部" 的员工信息
func getTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	// 不同点 2: 便捷查询方法
	// - database/sql: 提供基础查询方法如 Query、QueryRow，单条记录查询需手动处理 ErrNoRows 错误。
	// - sqlx: 提供 Get 和 Select 等便捷方法，Select 用于查询多条记录，自动处理 ErrNoRows 错误。
	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return employees, nil
}

// getHighestPaidEmployee 查询工资最高的员工信息
func getHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	// 不同点 2: 便捷查询方法
	// - database/sql: 单条记录查询需手动处理 ErrNoRows 错误。
	// - sqlx: Get 方法用于查询单条记录，自动处理 ErrNoRows 错误。
	err := db.Get(&employee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return Employee{}, nil
		}
		return Employee{}, err
	}
	return employee, nil
}

func main() {
	db, err := getDBConnection()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 查询所有部门为 "技术部" 的员工信息
	techEmployees, err := getTechEmployees(db)
	if err != nil {
		fmt.Println("查询技术部员工信息失败:", err)
	} else {
		fmt.Println("技术部员工信息:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	}

	// 查询工资最高的员工信息
	highestPaidEmployee, err := getHighestPaidEmployee(db)
	if err != nil {
		fmt.Println("查询工资最高的员工信息失败:", err)
	} else {
		fmt.Printf("工资最高的员工信息: ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n", highestPaidEmployee.ID, highestPaidEmployee.Name, highestPaidEmployee.Department, highestPaidEmployee.Salary)
	}

	// 不同点 3: 命名参数支持
	// - database/sql: 不支持命名参数，只能使用占位符 ?，复杂查询中参数顺序易混淆。
	// - sqlx: 支持命名参数，提高代码可读性和可维护性。示例如下：
	// _, err := db.NamedExec("INSERT INTO employees (name, department, salary) VALUES (:name, :department, :salary)",
	//     map[string]interface{}{
	//         "name":       "张三",
	//         "department": "技术部",
	//         "salary":     8000,
	//     })
}
