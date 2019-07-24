package balance

type Other struct{

}

func init(){
	RegisterBalancer("other", &Other{})
}

func(other *Other)DoBalance(insts []*Instance)(inst *Instance, err error){
	inst = insts[0]
	return 
}