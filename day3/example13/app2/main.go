package main

import (
	"fmt"
)
//完数
func isPerfectNumber(n int) bool {
	var factors []int
	sum:=0
	for i := 1; i <n; i++ {
		if n%i==0 {
			factors = append(factors, i)
			sum+=i
		}
	}
	
	return sum==n
}

func main(){
	fmt.Println("1000以内的完数有：")
	for i := 2; i < 1000; i++ {
		if isPerfectNumber(i){
			fmt.Println(i)
		}
	}

}