package cmd

import (
	"go_dev/day9/7_problems/model"
	"go_dev/day9/7_problems/handler"
)

const (
	Login = "login"
	Register = "register"
)

type Cmd interface{
	Execute (*Request) *Response
}

type LoginCmd struct{
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (loginCmd *LoginCmd) Execute(request *Request) *Response{
	request.DecodeData(loginCmd)
	u,err:=handler.Login(loginCmd.UserName, loginCmd.Password)
	if err!=nil{
		return NewResponse(false, err.Error(), "", "")
	}
	res:=NewResponse(true, "", "", "")
	res.EncodeBody(u)
	return res
}

type RegisterCmd struct{
	user *model.User
}

func (registerCmd *RegisterCmd) Execute(request *Request) *Response{
	request.DecodeData(registerCmd)
	user,err:=handler.Register(registerCmd.user)
	if err!=nil{
		return NewResponse(false, err.Error(), "", "")
	}
	res:=NewResponse(true, "", "", "")
	res.EncodeBody(user)
	return res
}