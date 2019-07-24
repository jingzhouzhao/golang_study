package linked

type LinkedNode struct{
	Data interface{}
	Next *LinkedNode
}

type LinkedList struct{
	head *LinkedNode
	tail *LinkedNode
}

func (p *LinkedList) InsertHead(data interface{}){
	node:=&LinkedNode{
		Data:data,
	}
	if p.head == p.tail && p.head==nil{
		p.head = node
		p.tail = node
		return 
	}
	node.Next = p.head
	p.head = node
}


func (p *LinkedList) InsertTail(data interface{}){
	node:=&LinkedNode{
		Data:data,
	}
	if p.head == p.tail && p.head==nil{
		p.head = node
		p.tail = node
		return 
	}
	p.tail.Next = node
	p.tail = node
}


func (p *LinkedList) GetHead() (head *LinkedNode){
	head = p.head
	return
}