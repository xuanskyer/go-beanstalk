package ModelFields

import "time"

//数据表模型字段
type TestModel struct {
	Id         int     `gorm:"AUTO_INCREMENT"`
	Name       string
	Desc       string
	Title      string
	Enum_a      string
	Status		int
	Created_at time.Time
}


func (TestModel) TableName() string {
	return "test"
}