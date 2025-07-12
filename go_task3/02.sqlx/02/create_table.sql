-- 创建 books 表
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- 插入示例数据
INSERT INTO books (title, author, price)
VALUES 
('Go 语言实战', '张三', 80.00),
('Python 入门教程', '李四', 60.00),
('Java 编程思想', '王五', 90.00),
('数据库原理', '赵六', 40.00);