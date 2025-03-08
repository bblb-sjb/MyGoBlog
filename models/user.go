package models

import "time"

type User struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Password string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreatAt  time.Time `json:"creatAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
