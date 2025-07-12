package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Student 定义学生结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// 数据库连接信息，需要根据实际情况修改
const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "test"
)

// getDBConnection 获取数据库连接
func getDBConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	db, err := sql.Open("mysql", dsn)
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

// insertStudent 插入新学生记录
func insertStudent(db *sql.DB, name string, age int, grade string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, age, grade)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// queryStudents 查询年龄大于 18 岁的学生
func queryStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT id, name, age, grade FROM students WHERE age > 18")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

// updateStudentGrade 更新学生年级
func updateStudentGrade(db *sql.DB, name string, newGrade string) (int64, error) {
	stmt, err := db.Prepare("UPDATE students SET grade = ? WHERE name = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newGrade, name)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// deleteStudents 删除年龄小于 15 岁的学生
func deleteStudents(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM students WHERE age < 15")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func main() {
	db, err := getDBConnection()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 插入新记录
	insertID, err := insertStudent(db, "张三", 20, "三年级")
	if err != nil {
		fmt.Println("插入记录失败:", err)
	} else {
		fmt.Println("插入记录成功，ID 为:", insertID)
	}

	// 查询年龄大于 18 岁的学生
	students, err := queryStudents(db)
	if err != nil {
		fmt.Println("查询记录失败:", err)
	} else {
		fmt.Println("年龄大于 18 岁的学生信息:")
		for _, student := range students {
			fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n", student.ID, student.Name, student.Age, student.Grade)
		}
	}

	// 更新学生年级
	rowsAffected, err := updateStudentGrade(db, "张三", "四年级")
	if err != nil {
		fmt.Println("更新记录失败:", err)
	} else {
		fmt.Println("更新记录成功，受影响的行数:", rowsAffected)
	}

	// 删除年龄小于 15 岁的学生
	rowsAffected, err = deleteStudents(db)
	if err != nil {
		fmt.Println("删除记录失败:", err)
	} else {
		fmt.Println("删除记录成功，受影响的行数:", rowsAffected)
	}
}
