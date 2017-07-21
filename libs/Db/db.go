package Db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"go-beanstalk/libs/Db/ModelFields"
)

func Test() {
	mysql, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer mysql.Close()

	// Migrate the schema
	mysql.AutoMigrate(&ModelFields.TestModel{})

	// Read
	var info ModelFields.TestModel
	var list []ModelFields.TestModel
	mysql.Find(&info) // find product with id 1

	fmt.Println(info)
	//fmt.Println(info.Id)
	//fmt.Println(info.Enum_a)
	//fmt.Println(info.Status)
	//fmt.Println(info.Created_at)

	mysql.Find(&list, "title = ?", "title1") // find product with code l1212
	fmt.Println(list)
	for k, v := range list {
		fmt.Println(k, v)
	}

	// Delete - delete product
	//mysql.Delete(&product)
}
