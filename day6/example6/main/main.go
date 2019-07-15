package main

import (
	"time"
	"fmt"
	"go_dev/day6/example6/balance"
	"math/rand"
	"os"
)

func main(){
	var insts [] *balance.Instance
	for i := 0; i < 10; i++ {
		var inst *balance.Instance
		inst = inst.NewInstance(fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)), 8080)
		insts = append(insts, inst)
	}
	
	config:="random"
	if len(os.Args)>1{
		config = os.Args[1]
	}

	for  {
		inst,err:=balance.DoBalance(config, insts)
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
	
}