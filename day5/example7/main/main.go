package main

import (
	t "go_dev/day5/example7/test"
	"fmt"
)

type Student struct{
	name string
	age int
	score float32
}

//声明一个方法，与其它语言重大区别
//(p Student) 表示这个方法作用于Student类型的p变量
func (p Student) init(name string,age int,score float32){
	p.name = name
	p.age = age
	p.score = score
	//此处p只是外面变量的一个副本，此处改变了p的字段，外面的变量并不会改变，正确做法是使用指针，如下
}

//go中方法/函数不支持重载，所以此处方法名不能同名
func (p *Student) init1(name string ,age int ,score float32){
	p.name = name
	p.age = age
	p.score = score
	//通常正确的赋值写法应该是(*p).name=xxx，但是go，结构体类型的可以直接简写为使用变量名。如果是其他类型还是需要(*int变量=10)
}

type integer int

func (p *integer) set(val integer){
	//非结构体还是得使用*p=val来赋值
	//p = val
	*p = val
}

func main(){

	var stu Student 
	stu.init("副本", 18, 11.1)
	fmt.Println("stu并不会被改变",stu)

	//通常写法应该是：(&stu).init1("指针", 19, 22.22)，但是go type 变量可以简写，所以直接stu.init1也是传递的指针
	stu.init1("指针", 19, 22.22)
	fmt.Println("stu会被改变",stu)

	var i integer
	//(&i).set(10) 可以简写：
	i.set(10)
	fmt.Println("i会被改变，因为传递的是指针",i)

	t.Test()
}