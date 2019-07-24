package main

import (
	"fmt"
	"os"
)

var (
	input = "zhangsan 18 99.99"
	format = "%s %d %f"
)

type Student struct{
	name string
	age int
	score float32
}

func main(){
	//输入内容到文件
	file,err:=os.OpenFile("D:/test.log", os.O_CREATE|os.O_RDWR, 0664)
	if err!=nil{
		fmt.Println("open file err:",err)
	}
	defer file.Close()
	//fmt.F开头的表示从文件/Writer/Reader读写
	fmt.Fprintln(file, "15 line ,hahahhaooooo fmt output test")
	//fmt.Fscanf(r, format, a)
	//fmt.S开头的表示从终端读写
	//fmt.Scan(a)
	//fmt.Scanln(a)
	//fmt.Scanf(format, a)
	//fmt.Sprintf(format, a)
	//fmt.Ss开头的表示从字符串读
	var stu Student
	fmt.Sscanf(input, format, &stu.name,&stu.age,&stu.score)
	fmt.Println(stu)


}