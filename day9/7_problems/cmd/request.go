package cmd

import (
	"fmt"
	"encoding/json"
)

type Request struct{
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}

func (req *Request) DecodeData(cmd Cmd){
	err:=json.Unmarshal([]byte(req.Data), cmd)
	if err!=nil{
		fmt.Println("DecodeData Error:",err)
	}
}