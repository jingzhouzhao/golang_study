package main

import (
	"fmt"
)

//在golang中，接口就是一组方法定义，一个结构体实现一个接口，只需要实现这个接口所有的方法即可。不用显示指定实现。例如java中的implement
type Animal interface{
	Crying()
	Running()
	//golang中，接口不能定义字段
	//name string
}

type Bird interface{
	Flying()
}

type Dog struct{
}

func (dog Dog) Crying(){
	fmt.Println("狗在叫，汪汪汪")
}

func (dog Dog) Running(){
	fmt.Println("狗在跑，~~~~")
}

type Cat struct{
}

func (cat Cat) Crying(){
	fmt.Println("猫在叫，喵喵喵")
}

func (cat Cat) Running(){
	fmt.Println("猫在跑，biu biu biu")
}

type Duck struct{

}

func (duck Duck) Crying(){
	fmt.Println("鸭子在叫，嘎嘎嘎")
}

func (duck Duck) Flying(){
	fmt.Print("鸭子在飞！！！！")
}

func main(){
	var animal Animal
	var dog Dog
	animal = dog
	animal.Crying()
	animal.Running()

	var cat Cat
	animal = cat
	animal.Crying()
	animal.Running()

	var duck Duck
	//由于duck没有实现Animal中的所有方法，所以duck不算实现了Animal接口，于是无法赋值
	//animal = duck
	duck.Crying()
	var bird Bird
	bird = duck
	bird.Flying()

}