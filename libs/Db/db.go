package Db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"go-beanstalk/libs/Db/ModelFields"
	"go-beanstalk/libs/Config"
)

var (
	dsn string = ""
)

func buildDsn() {
	test_config := Config.Get("db.test")
	database := test_config["dbname"]
	address := ""
	if "" != test_config["host"] {
		address = test_config["host"]
	}
	if "" != test_config["port"] {
		address = address + ":" + test_config["port"]
	}
	dsn = test_config["username"] + ":" + test_config["password"] + "@" + test_config["protocol"] + "(" + address + ")/" + database + "?charset=" + test_config["charset"] + "&parseTime=" + test_config["parseTime"] + "&loc=" + test_config["loc"]
}

func Test() {
	buildDsn()
	//mysql, err := gorm.Open("mysql", "root:root@tcp()/test?charset=utf8&parseTime=True&loc=Local")
	mysql, err := gorm.Open("mysql", dsn)
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
