package dao

import (
	"database/sql"
	"log"
)

func GetUserNameById(uid int) string {
	var name string
	err := DB.QueryRow("SELECT user_name FROM blog_user WHERE uid = ?", uid).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("未找到用户ID %d 的用户名", uid)
			return ""
		}
		log.Printf("查询用户名异常：%v", err)
		return ""
	}
	return name
}
