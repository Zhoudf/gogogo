-- 创建 employees 表
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(255) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL
);

-- 插入示例数据
INSERT INTO employees (name, department, salary)
VALUES 
('张三', '技术部', 8000.00),
('李四', '市场部', 6000.00),
('王五', '技术部', 9000.00);