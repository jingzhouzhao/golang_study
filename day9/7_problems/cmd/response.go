package cmd

import (
	"fmt"
	"encoding/json"
)

type Response struct{
	Success bool `json:"success"`
	Error string `json:"error"`
	Body string `json:"data"`
	Extra string `json:"extra"`
}

func NewResponse(success bool,err string,body,extra string) *Response {
	return &Response{
		Success:success,
		Error:err,
		Body:body,
		Extra:extra,
	}
}

func (res *Response) Bytes() []byte{
	data,err:=json.Marshal(res)
	if err!=nil{
		return []byte("response serialize error")
	}
	return data
}

func (res *Response) EncodeBody(body interface{}){
	data,err:=json.Marshal(body)
	if(err!=nil){
		res.Success = false
		fmt.Println("Encode body Error:",err)
		return
	}
	res.Body = string(data)
}

