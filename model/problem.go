/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-18 14:07:37
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-18 14:37:35
 * @FilePath: \Gorm\model\problem.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Cid      string `gorm:"column:cid;type:varchar(255);" json:"cid"`
	Title    string `gorm:"column:title;type:varchar(255);" json:"title"`
	Content  string `gorm:"column:content;type:text;" json:"content"`
}

func (table *Problem) TableName() string {
	return "problem"
}
