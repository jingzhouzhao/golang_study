package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	//中文不友好 因为一个中文等于三个字节，如果用下标的话。就不正确了
	// length := len(str)
	// for i := 0; i < length; i++ {
	// 	if i == length/2 {
	// 		break
	// 	}
	// 	last:=length-i-1
	// 	if str[i]!=str[last]{
	// 		return false
	// 	}
	// }


	//中文友好 rune代表一个字符
	t := []rune(str)
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Println("请输入：")
	fmt.Scanf("%s\n", &input)
	fmt.Println("你输入的是：", input)
	if isPalindrome(input) {
		fmt.Printf("%s是回文字符串", input)
	} else {
		fmt.Printf("%s不是回文字符串", input)
	}

}
