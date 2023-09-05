package services

import (
	"dudu/models"
	"dudu/util"
	"encoding/json"
	"errors"
	"log"
	// jwt "github.com/dgrijalva/jwt-go"
)

// 定义userService
type UserService struct {
}

func (*UserService) CreateNewUser(body []byte) (int64, error) {
	req := models.User{}
	json.Unmarshal(body, &req)

	user := models.User{}
	res, _ := req.HasUser(req.Telephone)

	if res.Id != 0 {
		res.PassWord = "nil"
		return int64(res.Id), errors.New("用户已存在")
	}

	id, err := user.Insert(req)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int64(id), nil
}

func (*UserService) UserLogin(email, telephone, password string) (int16, error) {
	if email == "" && telephone == "" {
		return 0, errors.New("请提供手机号或邮箱实现登录")
	}

	if password == "" {
		return 0, errors.New("~请提供密码！")
	}

	localUser := models.User{}

	if email != "" {
		req := models.User{}
		req.Email = email
		res := localUser.FindUserByEmail(req)
		localUser = res
	} else {
		user, err := localUser.HasUser(string(telephone))
		if err != nil {
			return 0, errors.New("没有找到相关用户，请去注册~")
		}

		localUser = user
		localUser.Id = user.Id
	}

	if localUser.Id == 0 {
		return 0, errors.New("没有找到相关用户，请去注册~")
	}

	valid := util.ComparePwd(string(localUser.PassWord), string(password))
	if !valid {
		return 0, errors.New("密码错误~")
	}

	return int16(localUser.Id), nil
}

func init() {

}
