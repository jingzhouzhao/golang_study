package model

import (
	"sync/atomic"
	"time"
	"fmt"
)

type BorrowBook interface{
	BorrowBook(book Book)
}

type Student struct{
	Name string
	Grade string
	IDCard string
	Gender rune
	Books []*Book
}

func(stu *Student)Init(name string,grade string,idCard string,gender rune){
	stu.Name = name
	stu.Grade = grade
	stu.IDCard = idCard
	stu.Gender = gender
}

func (stu *Student) BorrowBook(book *Book){
	stu.Books = append(stu.Books, book)
}
func (stu *Student) String() string{
	var str string = fmt.Sprintf("姓名：%s,身份证号: %s 已借书籍：\n",stu.Name,stu.IDCard)
	for i:=0;i<len(stu.Books); i++ {
		str+=fmt.Sprintf("\t %d.《%s》 作者：%s\n",i+1, stu.Books[i].Name,stu.Books[i].Author)
	}
	return str
}

type Record interface{
	Record(stu Student)
}


type Book struct{
	Name string
	Number int32
	Author string
	PulishDate time.Time
	Students []*Student
}

func (book *Book) Init(name string ,number int32 ,author string,pulishDate time.Time){
	book.Name = name
	book.Number = number
	book.Author = author
	book.PulishDate = pulishDate
}

func (book *Book) Record(stu *Student){
	book.Students = append(book.Students, stu)
	atomic.AddInt32(&book.Number, -1)
}

func (book *Book) String() string{
	return fmt.Sprintf("书名：《%s》 作者：%s 出版日期：%s 副本数：%d\n",book.Name,book.Author,book.PulishDate.Format("2006-01-02"),book.Number)
}
