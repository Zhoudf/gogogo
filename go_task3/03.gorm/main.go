package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name      string
	PostCount int
	Posts     []Post `gorm:"foreignKey:UserID"`
}

// TableName 自定义 User 结构体对应的表名
func (User) TableName() string {
	return "gorm_users"
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title        string
	Content      string
	UserID       uint
	User         User
	CommentCount int
	// CommentStatus  string
	Comments []Comment `gorm:"foreignKey:PostID"`
}

func (Post) TableName() string {
	return "gorm_posts"
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	Post    Post
}

func (Comment) TableName() string {
	return "gorm_comments"
}

// 2
// getUserPostsWithComments 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// getPostWithMostComments 查询评论数量最多的文章信息
func getPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post
	err := db.Model(&Post{}).Order("comment_count DESC").First(&post).Error
	return post, err
}

// BeforeCreate 在文章创建时更新用户的文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return err
	}
	user.PostCount++
	return tx.Save(&user).Error
}

// AfterDelete 在评论删除时检查文章的评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var post Post
	if err := tx.Preload("Comments").First(&post, c.PostID).Error; err != nil {
		return err
	}
	post.CommentCount = len(post.Comments)
	// if post.CommentCount == 0 {
	// 	post.CommentStatus = "无评论"
	// }
	return tx.Save(&post).Error
}

func main() {
	// 数据库连接信息，需要根据实际情况修改
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// // 1.模型定义
	// // 创建模型对应的数据库表
	// db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// // 创建用户
	// user := User{
	// 	Name:      "张三",
	// 	PostCount: 0,
	// }
	// if err := db.Create(&user).Error; err != nil {
	// 	fmt.Printf("创建用户失败: %v\n", err)
	// 	return
	// }
	// fmt.Printf("用户创建成功，ID: %d\n", user.ID)

	// // 创建文章
	// post := Post{
	// 	Title:        "第一篇文章",
	// 	Content:      "这是张三发布的第一篇文章内容。",
	// 	UserID:       user.ID,
	// 	CommentCount: 0,
	// }
	// if err := db.Create(&post).Error; err != nil {
	// 	fmt.Printf("创建文章失败: %v\n", err)
	// 	return
	// }
	// fmt.Printf("文章创建成功，ID: %d\n", post.ID)

	// // 创建评论
	// comment := Comment{
	// 	Content: "写得很不错！",
	// 	PostID:  post.ID,
	// }
	// if err := db.Create(&comment).Error; err != nil {
	// 	fmt.Printf("创建评论失败: %v\n", err)
	// 	return
	// }
	// fmt.Printf("评论创建成功，ID: %d\n", comment.ID)

	// // 更新文章评论数量
	// if err := db.Model(&post).Update("comment_count", 1).Error; err != nil {
	// 	fmt.Printf("更新文章评论数量失败: %v\n", err)
	// 	return
	// }
	// fmt.Println("文章评论数量更新成功")

	// // 2.关联查询
	// // 查询某个用户发布的所有文章及其对应的评论信息
	// userID := uint(1)
	// posts, err := getUserPostsWithComments(db, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("用户 ID 为 %d 的所有文章及其评论信息:\n", userID)
	// for _, post := range posts {
	// 	fmt.Printf("文章 ID: %d, 标题: %s\n", post.ID, post.Title)
	// 	fmt.Println("评论信息:")
	// 	for _, comment := range post.Comments {
	// 		fmt.Printf("  评论 ID: %d, 内容: %s\n", comment.ID, comment.Content)
	// 	}
	// }

	// // 查询评论数量最多的文章信息
	// post, err := getPostWithMostComments(db)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\n评论数量最多的文章信息:\n")
	// fmt.Printf("文章 ID: %d, 标题: %s, 评论数量: %d\n", post.ID, post.Title, post.CommentCount)
	// fmt.Println("评论信息:")
	// for _, comment := range post.Comments {
	// 	fmt.Printf("  评论 ID: %d, 内容: %s\n", comment.ID, comment.Content)
	// }

	// 	1. 创建相关钩子
	// BeforeCreate(tx *gorm.DB) error：在记录创建到数据库之前触发。
	// AfterCreate(tx *gorm.DB) error：在记录成功创建到数据库之后触发。
	// 2. 更新相关钩子
	// BeforeUpdate(tx *gorm.DB) error：在记录更新操作执行之前触发。
	// AfterUpdate(tx *gorm.DB) error：在记录成功更新之后触发。
	// 3. 删除相关钩子
	// BeforeDelete(tx *gorm.DB) error：在记录删除操作执行之前触发。
	// AfterDelete(tx *gorm.DB) error：在记录成功删除之后触发。
	// 4. 查询相关钩子
	// BeforeFind(tx *gorm.DB) error：在查询操作执行之前触发。
	// AfterFind(tx *gorm.DB) error：在查询操作完成之后触发。
	// 5. 保存相关钩子
	// BeforeSave(tx *gorm.DB) error：在执行 Save 操作（包含创建和更新）之前触发。
	// AfterSave(tx *gorm.DB) error：在执行 Save 操作成功之后触发。
	// 6. 验证相关钩子
	// BeforeValidate(tx *gorm.DB) error：在执行验证操作之前触发，可用于预处理数据。
	// AfterValidate(tx *gorm.DB) error：在执行验证操作之后触发。

	// 3.钩子函数
	// 创建一个用户
	user := User{
		Name:      "TestUser",
		PostCount: 0,
	}
	db.Create(&user)

	// 创建一篇文章，触发 Post 的 BeforeCreate 钩子函数
	post := Post{
		Title:   "Test Post",
		Content: "This is a test post.",
		UserID:  user.ID,
	}
	db.Create(&post)

	// 打印用户的文章数量，验证钩子函数是否生效
	var updatedUser User
	db.First(&updatedUser, user.ID)
	fmt.Printf("用户 %s 的文章数量: %d\n", updatedUser.Name, updatedUser.PostCount)

	// 创建一条评论
	comment := Comment{
		Content: "Great post!",
		PostID:  post.ID,
	}
	db.Create(&comment)

	// 删除评论，触发 Comment 的 AfterDelete 钩子函数
	db.Delete(&comment)

	// 打印文章的评论状态，验证钩子函数是否生效
	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("文章 %s 的评论数量: %d\n", updatedPost.Title, updatedPost.CommentCount)

}
