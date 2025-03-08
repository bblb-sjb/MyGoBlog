package dao

import (
	"database/sql"
	"errors"
	"log"
	"myWeb/models"
)

func GetUserNameById(uid int) string {
	var name string
	err := DB.QueryRow("SELECT user_name FROM blog_user WHERE uid = ?", uid).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("未找到用户ID %d 的用户名", uid)
			return ""
		}
		log.Printf("查询用户名异常：%v", err)
		return ""
	}
	return name
}

func GetUser(userName string, passwd string) *models.User {
	var user models.User
	err := DB.QueryRow("SELECT uid, user_name, passwd, avatar, create_at, update_at FROM blog_user WHERE user_name = ? AND passwd = ?", userName, passwd).Scan(&user.Uid, &user.UserName, &user.Password, &user.Avatar, &user.CreatAt, &user.UpdateAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("未找到用户 %s", userName)
			return nil
		}
		log.Printf("查询用户异常：%v", err)
		return nil
	}
	return &user
}
