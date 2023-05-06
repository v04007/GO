package models

import (
	"encoding/json"
	"errors"
)

//定义结构体信息

type SingUpFrom struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

//UnmarshalJSON SignUpForm类型 自定义JSON反序列化的方法
func (s *SingUpFrom) UnmarshalJSON(data []byte) (err error) {
	require := struct {
		UserName        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &require)
	if err != nil {
		return
	} else if len(require.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(require.Password) == 0 {
		err = errors.New("缺少必填字段Password")
	} else if require.Password != require.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		s.UserName = require.UserName
		s.Password = require.Password
		s.ConfirmPassword = require.ConfirmPassword
	}
	return
}

type User struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
