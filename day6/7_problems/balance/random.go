package balance

import (
	"math/rand"
	"errors"
)

type RandomBalancer struct{

}

func init(){
	RegisterBalancer("random", &RandomBalancer{})
}

func (rb *RandomBalancer) DoBalance(insts []*Instance)(inst *Instance,err error) {
	len:=len(insts)
	if len == 0 {
		err = errors.New("No instance found")
		return
	}
	inst = insts[rand.Intn(len)]
	return
}