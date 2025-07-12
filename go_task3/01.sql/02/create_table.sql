-- 创建 accounts 表
CREATE TABLE accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00
);

-- 创建 transactions 表
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    from_account_id INT NOT NULL,
    to_account_id INT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    -- 添加外键约束，确保转出账户 ID 存在于 accounts 表
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),
    -- 添加外键约束，确保转入账户 ID 存在于 accounts 表
    FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

-- 插入账户数据，账户 1 余额 500 元，账户 2 余额 200 元
INSERT INTO accounts (balance) VALUES (500.00), (200.00);