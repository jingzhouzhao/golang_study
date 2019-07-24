package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"compress/gzip"
)

func main(){
	file,err:=os.Open("D:/test.log.gz")
	if err!=nil{
		fmt.Fprintf(os.Stdout, "Open File Error:%s", err)
		return 
	}
	defer file.Close()
	gReader,err:=gzip.NewReader(file)
	if err!=nil{
		fmt.Fprintf(os.Stderr, "Read Gzip File Error:%s", err)
		return 
	}
	defer gReader.Close()
	fmt.Printf("FileName:%s,Comment:%s,Extra:%s\n",gReader.Name,gReader.Comment,gReader.Extra )
	fmt.Println("读取结果：")
	bufReader:=bufio.NewReader(gReader)
	
	for{
		buf,_,err:=bufReader.ReadLine()
		
		if err!=nil{
			if err !=io.EOF{
				fmt.Fprintf(os.Stderr, "Read File Error:%s", err)
				return
			}
			break
		}
		fmt.Printf("%s\n",buf)
	}
	fmt.Println("读取完毕！")
}