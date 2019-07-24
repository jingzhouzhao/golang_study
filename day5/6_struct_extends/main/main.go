package main

import (
	"fmt"
)

type Car struct{
	name string
	length int
}

type Train struct{
	//匿名字段
	Car
	//同一个类型的匿名字段只能有一个
	//Car
	name string
	int
}

func main(){
	var t Train
	//匿名字段可以直接访问
	t.length=100
	//优先当前结构体中的字段
	t.name = "train"
	//匿名字段直接通过类型访问
	t.int=10
	//有冲突时通过类型Car.name ，没有冲突时可以简写为t.name	
	t.Car.name="car"
	fmt.Println(t)

}