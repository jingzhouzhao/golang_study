package main

import (
	"encoding/json"
	"bytes"
	"encoding/binary"
	"fmt"
	"go_dev/day9/example7/errors"
	"net"
	"go_dev/day9/example7/config"
	"go_dev/day9/example7/cmd"
)

func StartServer() (err error){
	fmt.Println("Starting Server...")
	listen,err:=net.Listen("tcp", config.ServerAddress)
	if err!=nil {
		err = errors.ListenError
		return
	}
	for{
		conn,err:=listen.Accept()
		if err!=nil {
			fmt.Println("accept error")
			continue
		}
		fmt.Println("Process Request...")
		go process(conn)
	}
}

func process(conn net.Conn){
	defer conn.Close()
	data:=readPackage(conn)
	if data != ""{
		req,err:=processBody(data)
		if err!=nil{
			conn.Write(cmd.NewResponse(false,err.Error(), "", "").Bytes())
			return
		}
		var c cmd.Cmd
		switch req.Cmd {
		case cmd.Login:
			c = &cmd.LoginCmd{}
		case cmd.Register:
			c = &cmd.RegisterCmd{}
		default:
			conn.Write(cmd.NewResponse(false,errors.UnknownCmd.Error(), "", "").Bytes())
			return
		}
		fmt.Println("Execute Cmd:",c,req)
		conn.Write(c.Execute(req).Bytes())
	}
}

func readPackage(conn net.Conn)(data string){
	buf:=make([]byte,8192)
	n,err:=conn.Read(buf[0:4])
	if n!=4{
		fmt.Println("Read header Error")
		conn.Write(cmd.NewResponse(false, errors.HeaderByteError.Error(), "test", "test2").Bytes())
		return 
	}
	var bodyLenght int32
	binary.Read(bytes.NewBuffer(buf[0:4]),binary.BigEndian,&bodyLenght)
	body,err:=conn.Read(buf[4:bodyLenght+1])
	if err !=nil{
		conn.Write(cmd.NewResponse(false,err.Error(), "", "").Bytes())
		return
	}
	data = string(body)
	return
}

func processBody(body string)(cmd *cmd.Request,err error){
	err = json.Unmarshal([]byte(body),cmd)
	return
}