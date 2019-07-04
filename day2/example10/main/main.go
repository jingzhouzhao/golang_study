package main

import (
	"fmt"
)

func main() {
	//切片，字符串截取
	str := "hello world"
	fmt.Printf("str=%s\n", str)
	substr := str[0:5]
	fmt.Printf("substr=%s\n", substr)
	//[startindex:length]
	//substr :=str[6:11] = substr :=str[6:]
	substr = str[6:]
	fmt.Printf("substr=%s\n", substr)

	result := reverse(str)
	fmt.Printf("result=%s\n", result)
	result = reverse1(result)
	fmt.Printf("result=%s\n", result)

}

func reverse(str string) string {
	length := len(str)
	var result string
	for i := 0; i < length; i++ {
		result += string(str[length-i-1])
	}
	return result
}

func reverse1(str string) string {
	length := len(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, str[length-i-1])
	}
	return string(result)
}
