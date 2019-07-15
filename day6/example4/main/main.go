package main

import (
	"fmt"
	"reflect"
)

func reflectTest(a interface{}){
	t:= reflect.TypeOf(a)
	v:=reflect.ValueOf(a)
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(v.Kind())
	fmt.Println(v.Type())

	//由接口转回原型
	//b:=v.Interface().(*Student)
	//fmt.Println(b)
}

type Student struct{
	Name string 
	Age int
}

func (stu *Student) SetName(name string){
	stu.Name = name
}

func main(){
	//reflectTest(1024)
	reflectTest(new(Student))
}