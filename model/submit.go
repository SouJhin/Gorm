/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-18 14:34:22
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-19 23:31:35
 * @FilePath: \Gorm\model\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AEt
 */
package model

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	Path            string `gorm:"column:path;type:varchar(255);" json:"path"`
}

func (table *Submit) TableName() string {
	return "submit"
}
