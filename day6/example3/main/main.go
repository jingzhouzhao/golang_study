package main

import(
	"go_dev/day6/example3/linked"
	"fmt"
)

func main(){
	//var linkedList *linked.LinkedList = new(linked.LinkedList)
	var linkedList linked.LinkedList
	for i := 0; i < 10; i++ {
		//linkedList.InsertHead(i)
		//linkedList.InsertTail(i)
		linkedList.InsertTail(fmt.Sprintf("node.%d", i))
	}
	//遍历链表
	p:=linkedList.GetHead()
	for p!=nil{
		fmt.Println(p.Data)
		p = p.Next
	}
}