package main

import (
	"strings"
	"fmt"
)

func isPalindrome(str string) bool{
	length:= len(str)
	var temp []byte
	for i := 0; i < length/2; i++ {
		temp = append(temp, str[length-i-1])
	}
	return strings.HasPrefix(str, string(temp))
}

func main(){
	var input string
	fmt.Println("请输入：")
	fmt.Scanf("%s\n", &input)
	fmt.Println("你输入的是：",input)
	if isPalindrome(input) {
		fmt.Printf("%s是回文字符串",input )
	}else{
		fmt.Printf("%s不是回文字符串",input )
	}
	
}