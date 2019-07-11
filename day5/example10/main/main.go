package main

import (
	"bufio"
	"errors"
	"fmt"
	"go_dev/day5/example10/model"
	"go_dev/day5/example10/enums/gender"
	"os"
	"strconv"
	"strings"
	"time"
)

var bookStore []*model.Book
var stuStore []*model.Student

/*
*初始化
 */
func init() {
	bookStore = make([]*model.Book, 0)
	stuStore = make([]*model.Student, 0)
}

func main() {
	fmt.Println("\n==============欢迎使用图书管理系统===============")
	for {
		fmt.Println("请选择您要使用的功能：")
		fmt.Println(`
		1.书籍录入
		2.书籍查询
		3.学生录入
		4.学生查询
		5.借书
		6.书籍管理
		`)
		var option int
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			addBook()
		case 2:
			searchBook()
		case 3:
			addStudent()
		case 4:
			searchStudent()
		case 5:
			borrowBooks()
		case 6:
			showBookRecord()
		default:
			fmt.Println("未知选项")

		}

	}
}

type formatFunc func() error

func formatInput(f formatFunc) {
	var err error = errors.New("")
	for err != nil {
		err = f()
		if err != nil {
			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()
		}
	}
}

func input(tip string) string {
	fmt.Println(tip)
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("输入错误")
		return ""
	}
	return string(input)
}

func addBook() {
	var book model.Book
	book.Name = input("请输入书籍名称:")
	book.Author = input("请输入作者名称：")
	var pulishDate time.Time
	formatInput(func() error {
		var err error
		pulishDate, err = time.Parse("2006-01-02", input("请输入出版日期(yyyy-MM-dd)："))
		if err != nil {
			return errors.New("请输入正确日期")
		}
		return err
	})
	book.PulishDate = pulishDate
	var number int
	formatInput(func() error {
		var err error
		number, err = strconv.Atoi(input("请输入书籍数量:"))
		if err != nil {
			return errors.New("请输入正确的数量")
		}
		return err
	})
	book.Number = int32(number)
	bookStore = append(bookStore, &book)
	fmt.Println("书籍录入完成！")
	fmt.Println()
}

func searchBook() {
	var keyword string = input("请输入关键字：")
	fmt.Println("查询结果：")
	var index int = 1
	for _, v := range bookStore {
		//??代码很长如何换行
		if strings.Contains(v.Name, keyword) || strings.Contains(v.Author, keyword) || strings.Contains(v.PulishDate.Format("2006-01-02"), keyword) {
			fmt.Printf("%d.%s",index, v)
			index++
		}
	}
	fmt.Println()
}

func addStudent(){
	var stu model.Student 
	stu.Name = input("请输入学生姓名：")
	stu.Grade = input("请输入学生年级")
	formatInput(func() error{
		var err error
		in:=input("请选择学生性别(0-女,1-男)：")
		switch in {
		case "0":
			stu.Gender = gender.Female
		case "1":
			stu.Gender = gender.Male
		default:
			return errors.New("您的选择有误！请重新选择！")
		}
		return err
	})
	stu.IDCard = input("请输入学生身份证号")
	stuStore = append(stuStore, &stu)
}

func searchStudent(){
	var keyword string = input("请输入关键字：")
	fmt.Println("查询结果：")
	var index int = 1
	for _, v := range stuStore {
		//??代码很长如何换行
		if strings.Contains(v.Name, keyword) || strings.Contains(v.Grade, keyword) || strings.Contains(v.IDCard, keyword) {
			fmt.Printf("%d.%s",index, v)
			index++
		}
	}
	fmt.Println()
}


func borrowBooks(){
	var stuName string = input("请输入姓名：")
	var idCard string = input("请输入身份证号码：")
	var stu *model.Student = getStudentByNameAndIDCard(stuName, idCard)
	if stu == nil {
		fmt.Println("学生不存在")
		return
	}
	fmt.Printf("学生：%s,年级：%s,身份证号：%s\n",stu.Name,stu.Grade,stu.IDCard)
	var books [] *model.Book
	formatInput(func()error{
		var err error
		var bookName = input("请输入需要借阅的书名：")
		books = getBooksByName(bookName)
		if books == nil{
			return errors.New("未查询到需要借阅的书籍，请重新输入！")
		}
		return err
	})
	var num int
	var sBook *model.Book
	formatInput(func()error{
		var err error
		for i,v:=range books{
			fmt.Printf("%d.《%s》作者：%s\n", i+1,v.Name,v.Author)
		}
		num,err = strconv.Atoi(input("请选择您需要借阅的书籍序号："))
		if err!=nil || num<1 || num>len(books){
			err = errors.New("请输入正确的序号！")
		}
		return err
	})
	sBook =  books[num-1]
	if sBook.Number<=0 {
		fmt.Println("该书已被借完")
		return
	}
	stu.BorrowBook(sBook)
	sBook.Record(stu)
	fmt.Println("借书完成！")
}

func getStudentByNameAndIDCard(name string , idCard string) *model.Student{
	var stu *model.Student
	for _,v :=range stuStore{
		if v.Name == name && v.IDCard == idCard {
			stu = v
			break
		}
	}
	return stu
}

func getBooksByName(bookName string) []*model.Book{
	var books []*model.Book
	for _,v:=range bookStore{
		if v.Name==bookName{
			//这里如果bookStore中的不是指针，这里用&v会有一个问题
			//v是一个变量，&v取得实际上是这个变量指向的地址，当循环结束时，&v始终是指向最后一个。
			books = append(books, v)
		}
	}
	
	return books
} 

func showBookRecord(){
	for i,v:=range bookStore{
		fmt.Printf("%d.%s 借阅记录：\n", i+1,v)
		for j,vv:=range v.Students{
			fmt.Printf("\t%d. %s %s %s\n", j+1,vv.Name,vv.Grade,vv.IDCard)
		}
	}
}