package main

import (
	"strconv"
	"fmt"
)
/**大数相加*/
func addBigNumber(str1,str2 string) string {
	len1:=len(str1)
	len2:=len(str2)
	length:=len1
	//补齐，方便计算
	if len2>len1 {
		length = len2
		str1 = fmt.Sprintf("%0*s", length,str1)
	}else{
		str2 = fmt.Sprintf("%0*s", length,str2)
	}
	var up bool = false
	var result []byte
	for index := 0; index < length; index++ {
		i:= length-index-1
		var a,b int
		if(index<len1){
			a=int(str1[i]-'0')
		}else{
			a=0
		}
		if(index<len2){
			b=int(str2[i]-'0')
		}else{
			b=0
		}

		r:=a+b
		if up{
			r+=1
		}
		if r > 9 {
			up = true
			r=r-10
		}else{
			up = false
		}
		result=append([]byte{byte(r+'0')},result...)
	}
	//最后一个进位
	if up{
		result=append([]byte{byte('1')},result...)
	}
	return string(result)
}

func main()  {
	// var str1 string
	// var str2 string
	// fmt.Println("请输入两个大数，以空格分隔：")
	// fmt.Scanf("%s%s\n", &str1,&str2)
	// fmt.Println("和为：",addBigNumber(str1,str2))
	//9223372036854775807
	// bi, _ := strconv.ParseInt("0111111111111111111111111111111111111111111111111111111111111111", 2, 64)
	// fmt.Printf("%d", bi)
}