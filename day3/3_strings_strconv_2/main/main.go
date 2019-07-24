package main

import (
	"strconv"
	"fmt"
	"strings"
)

func main()  {
	str:="ababcba";
	//替换，最后一个参数为次数，如果全部替换为-1
	str=strings.Replace(str, "ab", "ba", -1)
	fmt.Println(str)
	//统计一个字符串在另外一个字符串中出现的次数
	count:=strings.Count(str, "ba")
	fmt.Printf("ba出现了%d次\n",count)
	//重复
	str=strings.Repeat(str, 2)
	fmt.Printf("重复2次后的str是%s\n", str)
	//转大写
	str = strings.ToUpper(str)
	fmt.Println("str转换为大写:",str)

	str=" abc defa a23 sd,f"
	//去除空格
	str = strings.TrimSpace(str)
	fmt.Printf("去掉空格后的str:%s\n", str)
	str = strings.Trim(str, "a")
	fmt.Println("去掉首尾的a之后的str:",str)
	str = strings.TrimLeft(str, "b")
	fmt.Println("去掉左边的b之后的str：",str)
	var strs []string = strings.Fields(str)
	fmt.Println("按空格拆分后的结果：",strs)
	strs = strings.Split(str, ",")
	fmt.Println("按,拆分后的结果：",strs)
	str = strings.Join(strs, "-")
	fmt.Println("用-连接后的结果：",str)

	//strconv
	var a int = 123
	str = strconv.Itoa(a)
	fmt.Printf("int转换为string:%q\n",str)
	b,err := strconv.Atoi("adf")
	fmt.Printf("string转换为int:%d,err:%v\n", b,err)



}