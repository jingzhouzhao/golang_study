package balance

import (
	"fmt"
)

type BalanceMgr struct{
	balancer map[string]Balancer
}

var mgr = BalanceMgr {
	balancer:make(map[string]Balancer),
}

func (mgr *BalanceMgr) registerBalancer(balancerName string,b Balancer){
	mgr.balancer[balancerName] = b
}

func RegisterBalancer(balancerName string,b Balancer){
	mgr.registerBalancer(balancerName, b)
}

func DoBalance(balancerName string,insts []*Instance)(inst *Instance,err error){
	balancer,ok := mgr.balancer[balancerName]
	if !ok {
		err = fmt.Errorf("balancer[%s] not found", balancerName)
		return
	}
	fmt.Printf("use %s balancer\n",balancerName)
	inst,err = balancer.DoBalance(insts)
	return 
}