package balance

import (
	"errors"
)

type RoundRobinBalancer struct{
	curIndex int
}

func init(){
	RegisterBalancer("roundrobin", &RoundRobinBalancer{})
}

func (rrb *RoundRobinBalancer) DoBalance(insts []*Instance)(inst *Instance,err error) {
	len:=len(insts)
	if len == 0 {
		err = errors.New("No instance found")
		return
	}
	inst= insts[rrb.curIndex]
	rrb.curIndex = (rrb.curIndex+1) % len
	return 
}