package combine

import (
	"fmt"
)

func Combine(a int){
	for i:=0;i<=5;i++ {
		//fmt.Println(i,"+",a-i,"=",a)
		fmt.Printf("%d+%d=%d \n",i,a-i,a)
	}
}