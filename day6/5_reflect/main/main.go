package main

import (
	"fmt"
	"reflect"
)

type UnknownType struct{
	S1 string `json:"name" other:"oname"`
	i1 int
	s2 string
}

func (p *UnknownType) Set(s1 string,i1 int,s2 string){
	p.S1 = s1
	p.i1= i1
	p.s2 = s2
}

func Test(a interface{}){
	t:=reflect.TypeOf(a)

	fmt.Printf("a 是 %s 类型\n",t)
	//v:=reflect.ValueOf(t)
	kd:=t.Kind()
	fmt.Printf("a 是 %s 类别\n",kd)
	
	if t.Elem().Kind() != reflect.Struct{
		fmt.Println("a 不是一个 struct")
		return
	}
	//t = t.Elem() //因为t传递过来的是指针，指针是没有字段的。
	fmt.Printf("a 有 %d个字段\n",t.Elem().NumField())
	fmt.Printf("a 有 %d个方法\n",t.NumMethod())//这里是0个方法，因为Set方法是作用在指针上的
	fmt.Printf("a 的第0个字段是%v\n", t.Elem().Field(0).Name)
	fmt.Printf("a 的第0个字段的json tag是：%s\n",t.Elem().Field(0).Tag.Get("json"))
	fmt.Printf("a 的第0个字段的other tag是：%s\n",t.Elem().Field(0).Tag.Get("other"))
	//调用方法
	v:=reflect.ValueOf(a)
	v.Method(0).Call([]reflect.Value{reflect.ValueOf("_s1"),reflect.ValueOf(2),reflect.ValueOf("_s2")})
	fmt.Println(v)
}

func main(){
	var ut UnknownType = UnknownType{
		S1:"ss1",
		i1:1,
		s2:"ss2",
	}
	Test(&ut)
	fmt.Println(ut)
}