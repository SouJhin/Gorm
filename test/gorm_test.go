package test

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/models"
)

func TestGormTest(t *testing.T) {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&models.Food{})
	data := make([]*models.Food, 0)

	t1 := models.Test{1, "hahah"}
	db.Create(&t1)
	f1 := models.Food{
		"www.baidu.com",
		"f111",
		"hahahah",
		"2222",
		"miaoshu",
		time.Now(),
	}
	db.Create(&f1)
	err = db.Find(&data).Error
	//newPro := models.Problem{
	//	Identity:   "wocaocao",
	//	CategoryId: "123123",
	//	Title:      "titleleee",
	//	Content:    "asdasdasd",
	//	MaxRuntime: 11,
	//	MaxMem:     11,
	//}
	//
	//result := db.Create(&newPro) // 通过数据的指针来创建
	//fmt.Printf("result =====> 🚀🚀🚀 %v\n", result)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("data =====> 🚀🚀🚀 %v\n", data)
	for _, value := range data {
		fmt.Printf("value =====> 🚀🚀🚀 %v\n", value)
	}
}
