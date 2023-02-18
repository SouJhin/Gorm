/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-18 14:34:22
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-18 14:34:28
 * @FilePath: \Gorm\model\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AEt
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`
	Password string `gorm:"column:password;type:varchar(32);" json:"password"`
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Mail     string `gorm:"column:mail;type:varchar(100);" json:"mail"`
}

func (table *User) TableName() string {
	return "user"
}
