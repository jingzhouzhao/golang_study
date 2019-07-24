package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	//str:="123"
	str := "asdf"
	//string 转 int strconv.Atoi
	a, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can't convert %q to int \n", str)
	} else {
		fmt.Println(a)
	}
	//strings 包的使用
	var (
		url  string
		path string
	)
	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Printf("url=%s\n",url)
	fmt.Printf("path=%s\n",path)
}

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "https://")
	if !result {
		return fmt.Sprintf("https://%s", url)
	}
	return url
}

func pathProcess(path string) string{
	result := strings.HasSuffix(path, "/")
	if !result{
		return fmt.Sprintf("%s/", path)
	}
	return path
}