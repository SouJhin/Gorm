package models

import (
	"time"
)

type Food struct {
	//gorm.Model
	Image     string    `gorm:"column:image;type:varchar(255);" json:"identity"`
	Name      string    `gorm:"column:name;type:varchar(100);" json:"name"`
	ID        string    `gorm:"column:id;type:varchar(255);" json:"parent_id"`
	Price     string    `gorm:"column:price;type:int(255);" json:"price"`
	Describe  string    `gorm:"column:describe;type:varchar(255);" json:"describe"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
}
