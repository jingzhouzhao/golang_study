package test

import (
	"time"
	"testing"
)

func TestSave(t *testing.T){
	user:=&user{
		UserName:"张三",
		Email:"zhangsan@qq.com",
	}
	err:=user.Save()
	if err!=nil{
		t.Fatalf("test Save error:%s", err)
	}
	//成功可以不用t.logF
	t.Log("test Save success")
}

func TestLoad(t *testing.T){

	u:=&user{
		UserName:"张三",
		Email:"zhangsan@qq.com",
	}
	err:=u.Save()
	if err!=nil{
		t.Fatalf("test Save error:%s", err)
	}
	u1 := &user{}
	time.Sleep(10*time.Second)
	err = u1.Load()
	if err!=nil{
		t.Fatalf("test Load error:%s", err)
	}
	if u1.UserName != u.UserName{
		t.Fatalf("test Load error,UserName not equal")
	}
	if u1.Email != u.Email{
		t.Fatalf("test Load error,Email not equal")
	}
	
	t.Logf("test Load Success")

}