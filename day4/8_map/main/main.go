package main

import (
	"sort"
	"fmt"
)

func test(){
	//声明以及初始化map的几种方式
	var a map[string]string = map[string]string{"a":"author","b":"brother"}
	var b map[string]string = make(map[string]string)
	fmt.Println(a)
	b["a"]="adequate"
	b["b"]="before"
	fmt.Println(b)
	//删除map中的元素
	delete(b,"a")
	fmt.Println(b)
	//map的长度
	fmt.Println(len(a))
	//遍历map
	var c map[string]map[string]string
	c=make(map[string]map[string]string)
	c["key1"] = make(map[string]string)
	c["key1"]["a1"]="123"
	c["key1"]["a2"]="456"
	c["key1"]["a3"]="789"
	c["key2"] = make(map[string]string)
	c["key2"]["a1"]="000"
	c["key2"]["a2"]="111"
	c["key2"]["a3"]="222"

	for k,v:=range c{
		fmt.Println(k)
		for k1,v1 :=range v{
			fmt.Println("\t",k1,v1)
		}
	}

	//以上遍历map，key是无序的。应该通过有序的key列表来遍历map
	
	
}

func test2(a map[string]int){
	a["test"]=100
}

func test3(){
	fmt.Println("-----------------")
	var aa map[int]int
	aa= make(map[int]int)
	aa[1]=1
	aa[2]=2
	aa[3]=3
	aa[4]=4
	aa[5]=5
	for k,v:=range aa{
		fmt.Println(k,v)
	}
	fmt.Println("-------------------,right:")
	//以上为不正确的遍历方法，以下为正确的遍历方法
	var keys []int
	for k,_ :=range  aa{
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _,v:=range keys{
		fmt.Printf("%d-%d\n",v,aa[v])
	}
}
func main(){
	// test()
	// //map是引用传递（地址）
	// var a map[string]int=map[string]int{"test":10}
	// fmt.Println(a)
	// test2(a)
	// fmt.Println(a)
	//test3()
	var keys []int
	keys[0]=10
	keys = append(keys, 1)
}