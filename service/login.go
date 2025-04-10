package service

import (
	"errors"
	"myWeb/dao"
	"myWeb/models"
	"myWeb/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "bblb")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}
	uid := user.Uid
	//生成token  jwt技术进行生成 令牌  A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
