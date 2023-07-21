package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name       string `gorm:"column:name"`
	Department string `gorm:"column:department;de	fault:'develop'"`
}

func (User) TableName() string {
	return "backenddevelopment"
}

func Index(context *gin.Context) {
	context.String(200, "Hello world!")
}

func Json(c *gin.Context) {
	c.JSON(200, map[string]string{"2": "200"})
}

func creat(db *gorm.DB) {
	newRecord := User{
		Name: "John Doe",
	}
	results := db.Create(&newRecord)
	if results.Error != nil {
		panic("插入数据失败，error=" + results.Error.Error())
	}
	fmt.Printf("插入成功的记录：ID: %d, Name: %s, Department: %s\n", newRecord.ID, newRecord.Name, newRecord.Department)
}
func Delete(db *gorm.DB) {
	result := db.Where("id >= ? AND id <= ?", 8, 14).Delete(&User{})
	if result.Error != nil {
		fmt.Printf("范围批量删除失败，错误：%v\n", result.Error)
	} else {
		fmt.Printf("范围批量删除成功，受影响的记录数：%v\n", result.RowsAffected)
	}
}
func main() {
	dsn := "root:fangzx123456@tcp(127.0.0.1:3306)/relation?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败，erro=" + err.Error())
	}
	println(db)
	db.AutoMigrate(&User{})
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		panic("查询数据失败，error=" + result.Error.Error())
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Department: %s\n", user.ID, user.Name, user.Department)
	}
	//Delete(db)
	router := gin.Default() // 创建一个默认的路由
	router.LoadHTMLGlob("heml/*")
	router.GET("/index", Index)
	router.GET("/json", Json) // 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	router.StaticFS("/see", http.Dir("picture/Cansee"))
	router.StaticFile("/osee", "picture/img.png")
	router.GET("/", func(c *gin.Context) {
		var user User
		id := c.Query("id")
		result := db.First(&user, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{
					"message": "NOT FOUND",
				})
			} else {
				c.JSON(404, gin.H{
					"message": "查询错误：",
				})
			}
		} else {
			c.JSON(200, user)
		}

	})
	router.Run(":8087")
}
