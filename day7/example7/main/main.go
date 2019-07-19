package main

import (
	"io"
	"fmt"
	"os"
)

func CopyFile(dstName string ,srcName string)(written int64,err error){
	src,err:=os.Open(srcName)
	if err!=nil{
		fmt.Println("Open File Error:",err)
		return
	}
	defer src.Close()
	dst,err:=os.OpenFile(dstName, os.O_CREATE|os.O_RDWR, 0666)
	if err!=nil{
		fmt.Println("Writer File Error:",err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main(){
	CopyFile("D:/test.log.copy","D:/test.log")
}