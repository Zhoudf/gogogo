package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

// transferMoney 执行转账事务
func transferMoney(db *sql.DB, fromAccountID, toAccountID int, amount float64) error {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	//defer机制类似java:try-catch-finally
	defer func() {
		if p := recover(); p != nil {
			// 发生 panic，回滚事务
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// 发生错误，回滚事务
			tx.Rollback()
		} else {
			// 没有错误，提交事务
			err = tx.Commit()
		}
	}()

	// 检查账户 A 的余额是否足够
	var balance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromAccountID).Scan(&balance)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("账户余额不足")
	}

	// 从账户 A 扣除金额
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID)
	if err != nil {
		return err
	}

	// 向账户 B 增加金额
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccountID)
	if err != nil {
		return err
	}

	// 在 transactions 表中记录该笔转账信息
	_, err = tx.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)", fromAccountID, toAccountID, amount)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := getDBConnection()
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	fromAccountID := 1
	toAccountID := 2
	amount := 100.0

	err = transferMoney(db, fromAccountID, toAccountID, amount)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}
