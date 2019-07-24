package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func through(root *Student){
	if root == nil{
		return
	}
	//前序遍历
	// fmt.Println(root.Name)
	// through(root.left)
	// through(root.right)
	//中序遍历
	// through(root.left)
	// fmt.Println(root.Name)
	// through(root.right)
	//后序遍历
	through(root.left)
	through(root.right)
	fmt.Println(root.Name)
}

func main(){

	var stu *Student = new(Student)
	stu.Name="stu0"

	var stu1 *Student = new(Student)
	stu1.Name="stu1"
	var stu2 *Student = new(Student)
	stu2.Name="stu2"

	stu.left = stu1
	stu.right = stu2

	var stu3 *Student = new(Student)
	stu3.Name="stu3"

	stu1.left = stu3

	through(stu)




}