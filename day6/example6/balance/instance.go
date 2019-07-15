package balance

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func (inst *Instance) NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (inst *Instance) String()string{
	return inst.host+":"+strconv.Itoa(inst.port)
}

func (inst *Instance) GetHost() string {
	return inst.host
}

func (inst *Instance) GetPort() int {
	return inst.port
}
