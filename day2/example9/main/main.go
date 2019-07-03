package main

import (
	"time"
	"fmt"
	"math/rand"
)

func init(){
	//设置随机种子，不然每次随机结果一样，因为math/rand是伪随机
	seed := time.Now().UnixNano();
	fmt.Println("随机种子：",seed)
	rand.Seed(seed)
}

func main(){
	
	for i:=0;i<10;i++{
		fmt.Println(rand.Int())
		fmt.Println(rand.Int31n(100))
		fmt.Println(rand.Float32())
	}
}