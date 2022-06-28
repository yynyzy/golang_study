package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
gorm安装
操作MySQL需要安装两个包：


//安装MySQL驱动
go get -u gorm.io/driver/mysql
//安装gorm包
go get -u gorm.io/gorm
*/

/*
GORM定义一个 gorm.Model结构体，其包括字段ID、createdAt、UpdatedAt 、DeletedAt
ll gorm.Model的定义
````
type Model struct {
ID       uintl 'gorm:"primaryKey"
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt gorm.DeletedAt 'gorm:"index"
````
您可以将它嵌入到您的结构体中，以包含这几个字段，详情请参考嵌入结构体

*/
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//创建表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

//插入数据
func add(db *gorm.DB) {
	p := Product{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&p)
}

//查询数据
func find(db *gorm.DB) {
	var p Product
	// db.First(&p, 1)
	// fmt.Printf("p: %v\n", p)
	db.First(&p, "code=?", "1001")
	fmt.Printf("p: %v\n", p)

}

//更新
func update(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Model(&p).Update("price", "999")

	db.Model(&p).Updates(Product{Price: 1001, Code: "1002"}) //仅更新非零值字段
	// db.Model(&p).Updates(map[string]interface{}{"Price": 1003, "code": "1003"})

}

//删除
func delete(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Delete(&p, 1)
}
func main() {
	dsn := "root:yyn990902@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}
	// //延时关闭数据库连接
	// defer db.Close()

	// create(db)
	// add(db)
	// find(db)
	update(db)
}
