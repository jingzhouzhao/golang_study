package main

import (
	"os"
	"time"
	"fmt"
)

type PathError struct{
	Path string
	Op string 
	ErrorTime string
	Message string
}

func (err *PathError) Error() string{
	return fmt.Sprintf("Path:%s,Op:%s,ErrorTime:%s,Message:%s",err.Path,err.Op,err.ErrorTime,err.Message)
}

func Open(filename string) error{
	file,err:=os.Open(filename)
	if err!=nil{
		return &PathError{
			Path:filename,
			Message:err.Error(),
			ErrorTime:time.Now().Format("2006-01-02T15:04:05-07:00"),
			Op:"os.Open",
		}
	}
	fmt.Println(file)
	return nil
}

func main(){
	err:=Open("C:/123123123123")
	if v,ok:=err.(*PathError);ok{
		fmt.Fprintln(os.Stderr, v)
	}
}