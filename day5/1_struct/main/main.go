package main

import (
	"fmt"
)
//结构体是值类型
type Student struct{
	Name string
	Age int
	//通过首字母大小写来区分访问权限，小写字母包内可见
	score float32
}

func modifyScore(stu Student){
	stu.score = 100
}
func modifyScore2(stu *Student){
	//结构体的指针类型，可以不用(*stu).score，可以简写
	stu.score = 100
}


func main(){
	//结构体定义的几种方式
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 18
	stu.score=77.5
	fmt.Println(stu)
	//指针类型
	var stu1 *Student = &Student{
		Name : "lisi",
		Age : 19,
		score : 76,
	}
	// stu1.Name ="lisi"
	// stu1.Age = 19
	// stu.score = 76
	fmt.Println(stu1)
	var stu2 Student = Student{
		Name : "wangwu",
		Age : 20,
		//score : 77,
	}
	// stu2.Name = "wangwu"
	// stu2.Age = 20
	// stu2.score= 77
	fmt.Println(stu2)
	//指针类型
	var stu3 *Student = new(Student)
	stu3.Name = "zhaoliu"
	stu3.Age = 21
	stu3.score= 78
	fmt.Println(stu3)

	modifyScore(stu)
	fmt.Println(stu)
	modifyScore2(stu1)
	fmt.Println(stu1)
	modifyScore(stu2)
	fmt.Println(stu2)
	modifyScore2(stu3)
	fmt.Println(stu3)
	modifyScore2(&stu)
	fmt.Println(stu)

	var stu4 Student

	var stu5 Student
	//比较值
	fmt.Println(stu4==stu5)
}