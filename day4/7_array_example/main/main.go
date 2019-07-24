package main

import (
	"fmt"
	"sort"
)

func test(){
	var a =[...]int{3,5,1,3,6,8,7}
	sort.Ints(a[:])
	fmt.Println(a)

	var b = [...]string{"c","b","e","d","a"}
	sort.Strings(b[:])
	fmt.Println(b)

	var c = [...]float64{1.22,0.8,3.4,7.8}
	sort.Float64s(c[:])
	fmt.Println(c)


	index:=sort.SearchInts(a[:], 5)
	fmt.Printf("%d in a[] index:%d\n", 5,index)

}


func main()  {
	test()
	
}