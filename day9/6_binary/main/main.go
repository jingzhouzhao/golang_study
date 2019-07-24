package main

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

func main(){
	var x int32 = 123123123
	buffer:=bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, x)

	var buf []byte = buffer.Bytes()
	fmt.Println(buf)
	buffer=bytes.NewBuffer(buf[0:4])
	var v uint32
	//binary.BigEndian 大端
	err:=binary.Read(buffer, binary.BigEndian, &v)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(v)
}