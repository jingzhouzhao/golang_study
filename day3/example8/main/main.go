package main

import (
	"fmt"
)

func main()  {

	LABEL1:
	for index :=0 ; index <=5; index++{
		for j := 0; j <= 5; j++ {
			if j==4 {
				continue LABEL1
			}
			fmt.Printf("index:[%d],j:[%d]\n", index,j)
		}
	}
	
}