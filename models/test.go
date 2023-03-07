package models

type Test struct {
	//gorm.Model
	Id   int    `gorm:"column:id;type:int" json:"id"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}
