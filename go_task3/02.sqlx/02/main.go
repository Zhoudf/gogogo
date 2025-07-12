package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Book 定义书籍结构体，与 books 表字段对应
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// getDBConnection 获取数据库连接
func getDBConnection() (*sqlx.DB, error) {
	// 数据库连接信息，需要根据实际情况修改
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
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

// getBooksOver50 查询价格大于 50 元的书籍
func getBooksOver50(db *sqlx.DB) ([]Book, error) {
	var books []Book
	// 使用 sqlx 的 Select 方法执行查询，并将结果映射到 Book 结构体切片
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return books, nil
}

func main() {
	db, err := getDBConnection()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 查询价格大于 50 元的书籍
	books, err := getBooksOver50(db)
	if err != nil {
		fmt.Println("查询价格大于 50 元的书籍失败:", err)
	} else {
		fmt.Println("价格大于 50 元的书籍信息:")
		for _, book := range books {
			fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
		}
	}
}
