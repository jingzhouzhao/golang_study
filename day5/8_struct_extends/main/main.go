package main

import (
	"fmt"
)

type Animal struct{
	name string
	age int
}

type Dog struct{
	//Golang中通过匿名字段来实现继承。继承方法也是一样，如果匿名字段有方法，那么当前结构体自动继承
	Animal
}

type Cat struct{
	Animal
}
func (this Animal) Crying(){
	//因为不涉及到修改，所以此处Animal不用指针
	fmt.Printf("%s在叫...\n",this.name)
}

func (this Animal) Running(){
	//因为不涉及到修改，所以此处Animal不用指针
	fmt.Printf("%s在跑...\n",this.name)
}

//String()是一个接口，在打印时自动调用
func (this *Animal) String() string{
	str:=fmt.Sprintf("name:[%s],age:[%d]", this.name,this.age)
	return str
}

func main(){
	var dog Dog
	dog.name = "狗"
	dog.Crying()
	dog.Running()
	var cat Cat
	cat.name = "猫"
	cat.Crying()
	cat.Running()

	//因为实现String()方法的是Animal指针，所以这里要传地址才行。或者String 方法直接使用Animal类型
	//fmt.Println(cat)
	fmt.Println(&cat)

}