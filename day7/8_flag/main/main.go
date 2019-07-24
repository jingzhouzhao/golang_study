package main

import (
	"flag"
	// "os"
	 "fmt"
)

func main(){
	// fmt.Println("Args len:",len(os.Args))
	// for i,v:=range os.Args{
	// 	fmt.Printf("%d:%s\n", i,v)
	// }
	var configPath string
	var debug bool
	var threads int
	flag.StringVar(&configPath, "c", "", "配置文件路径")
	flag.BoolVar(&debug, "d", false, "是否开启debug")
	flag.IntVar(&threads, "t", 10, "线程数量，默认10个")
	flag.Parse()//一定要Parse才生效
	fmt.Printf("configPath:%s\n",configPath)
	fmt.Printf("debug:%t\n",debug)
	fmt.Printf("threads:%d\n",threads)
}