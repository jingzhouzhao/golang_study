package main

import (
	"fmt"
	"math/rand"
)

type Student struct{
	Name string
	Age int
	Score float32

	next *Student
}

func through(head *Student){
	p := head
	for p!=nil{
		fmt.Println(p)
		p = p.next
	}
}

func insertHead(head **Student){
	for i := 0; i < 10; i++ {
		var stu *Student = &Student{
			Name:fmt.Sprintf("stu%d", i),
			Age:rand.Intn(30),
			Score:rand.Float32()*100,
		}
		stu.next = *head
		*head = stu
	}
}

//链表尾部插入
func insertTail(head *Student){
	var tail = head
	for i := 0; i < 10; i++ {
		var stu *Student = &Student{
			Name:fmt.Sprintf("stu%d", i),
			Age:rand.Intn(30),
			Score:rand.Float32()*100,
		}
		tail.next = stu
		tail = stu
	}
}

func delNode(head *Student){
	p:=head
	prev := head
	for p!=nil {
		if p.Name=="stu6" {
			prev.next = p.next
			break;
		}
		prev = p
		p=p.next
	}
}

func insertNode(head *Student,newNode *Student){
	p:=head
	for p!=nil{
		if p.Name =="stu5" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p=p.next
	}
}

func main(){
	var stu *Student = &Student{
		Name:"张三",
		Age:18,
		Score:100,
	}
	//insertTail(&stu)
	insertHead(&stu)
	//through(stu)

	delNode(stu)
	through(stu)

	var newNode *Student = &Student{
		Name:"new stu",
		Age:18,
		Score:100,
	}

	insertNode(stu,newNode)
	through(stu)
}