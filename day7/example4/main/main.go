package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

type CharCounter struct {
	LetterCount int
	NumCount    int
	SpaceCount  int
	OtherCount  int
}

func main() {
	//打开一个文件
	file, err := os.Open("D:/test.log")
	if err != nil {
		fmt.Println("open file error:", err)
	}
	defer file.Close()
	var counter CharCounter
	var reader *bufio.Reader = bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		//结束
		if err == io.EOF {
			break
		}
		if err !=nil {
			fmt.Println("read failed:",err)
		}
		runes:=[]rune(str)
		for _,v:=range runes{
			switch {
			case v>='a' && v<='z':
				fallthrough
			case v>='A' && v<= 'Z':
				counter.LetterCount++
			case v>='0' && v<='9':
				counter.NumCount++
			case v==' '|| v=='\t':
				counter.SpaceCount++
			default:
				counter.OtherCount++
			}
		}
	}
	fmt.Printf("LetterCount:%d,NumCount:%d,SpaceCount:%d,OtherCount:%d",counter.LetterCount,counter.NumCount,counter.SpaceCount,counter.OtherCount)

}
