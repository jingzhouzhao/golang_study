package handler

import (
	"time"
	"go_dev/day9/7_problems/consts"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strings"
	"go_dev/day9/7_problems/model"
	"go_dev/day9/7_problems/errors"
	"go_dev/day9/7_problems/stores"
)

func Login(userName,password string) (user *model.User,err error){
	if strings.TrimSpace(userName)=="" || strings.TrimSpace(password) ==""{
		err = errors.ParameterError
		return
	}
	conn:=stores.GetConn()
	defer conn.Close()
	data,err:=redis.String(stores.HGet(conn, model.UserTable, userName))
	if err!=nil{
		if(err == redis.ErrNil){
			err = errors.UserNotExists
		}
		return
	}
	user = &model.User{}
	err = json.Unmarshal([]byte(data), user)
	if err!=nil{
		err = errors.DeserializeError
		return
	}
	if user.Password!=password{
		err = errors.PasswordError
		user = nil
		return
	}
	user.LastLoginTime = time.Now()
	user.Status = consts.Online
	user,err = save(user)
	return
}

func save(user *model.User) (u *model.User,err error){
	if user == nil{
		err = errors.ParameterError
		return
	}
	conn:=stores.GetConn()
	defer conn.Close()
	data,err:= json.Marshal(user)
	if err!=nil{
		err = errors.SerializeError
		return 
	}
	_,err = stores.HSet(conn, model.UserTable,user.UserName,string(data))
	u = user
	return
}

func Register(user *model.User)(u *model.User,err error){
	return save(user)
}