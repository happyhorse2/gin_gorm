package gormtest

import (
	"fmt"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Person :person结构体初始化
type Person struct {
	ID        uint
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
}

// User :user结构体初始化
type User struct {
	ID        uint
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
}

// Gormtest : go语言gorm框架使用
func Gormtest() {
	db, err := gorm.Open("mysql", "root:Helloworld930925@tcp(192.168.99.241:3306)/gin_gorm_db?charset=utf8")
	//db.SingularTable(true) //禁用表默认名，首字母小写
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	defer db.Close()
	// 检查模型`Person`表是否存在
	istableexit := db.HasTable("testpersons")
	fmt.Println(istableexit)
	//补充表的列，调整表结构和代码中结构体一致
	//db.AutoMigrate(&User{})
	//删除表
	//db.DropTable("users")
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	//添加记录
	// user := User{Firstname: "mayansong", Lastname: "mayansong"}
	// result := db.NewRecord(user)
	// db.Create(&user)
	// fmt.Println(result)
	// resultv := db.NewRecord(user)
	// fmt.Println(resultv)
	//db.Create(&user)

	//查询记录
	// user := User{}
	// db.Where("firstname = ?", "mayansong").First(&user)
	// fmt.Println(user)

	// //修改记录
	// user.Firstname = "mayansong"
	// db.Save(&user)

	//删除记录
	user := User{Firstname: "mayansong"}
	fmt.Println(user)
	db.Delete(&user) //注意如果没有主键的话，会删除满足条件的记录

}

// TableName : 重命名
func (Person) TableName() string {
	return "testpersons"
}
