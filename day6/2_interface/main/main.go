package main

import (
	"fmt"
)

type Reader interface{
	Read()
}

type Writer interface{
	Write()
}

type ReadWriter interface{
	//接口嵌套/继承
	Reader
	Writer
}

type File struct{
	Name string
}

func (f *File) Read(){
	fmt.Println("read data")
}
func (f *File) Write(){
	fmt.Println("write data")
}

func Test(rw ReadWriter){
	rw.Read()
	rw.Write()
}
func Test2(r Reader){
	//接口类型断言，类似Java中的 instanceof
	f,ok:=r.(*File)
	if !ok{
		fmt.Println("r cannot convert to *File")
		return
	}
	fmt.Println(f.Name)
}

func Test3(items ... interface{}){
	for i,v :=range items {
		switch v.(type) {
		case int,int16,int32,int64,int8:
			fmt.Printf("No.%d param is %T type,value is %v \n",i+1,v,v)
		case float32,float64:
			fmt.Printf("No.%d param is %T type,value is %v \n",i+1,v,v)
		case string:
			fmt.Printf("No.%d param is %T type,value is %v \n",i+1,v,v)
		case bool:
			fmt.Printf("No.%d param is %T type,value is %v \n",i+1,v,v)
		}
	}

}


func main(){
	var f File
	Test(&f)
	f.Name = "这是一个文件"
	Test2(&f)
	var r Reader
	Test2(r)

	Test3(1,2,2.2,8.8,9.9,float32(0.1),"asdf",true)
}