package test

import (
	"encoding/json"
	"io/ioutil"
)

type user struct{
	UserName string `json:"username"`
	Email string `json:"email"`
}


func (u *user) Save() error{
	data,err:=json.Marshal(u)
	if err !=nil{
		return err
	}
	err = ioutil.WriteFile("D:/user.dat", data, 0666)
	return err
}

func (d *user) Load()error{
	data,err:=ioutil.ReadFile("D:/user.dat")
	if err !=nil{
		return err
	}
	err = json.Unmarshal(data, d)
	return err
}
