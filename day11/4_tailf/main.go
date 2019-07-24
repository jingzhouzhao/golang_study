package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	filename := "C:\\Users\\zhaojz\\go\\src\\go_dev\\day11\\tailf\\a.log"
	t, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file error:", err)
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-t.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", filename)
			continue
		}
		//fmt.Printf("%#v\n",msg)
		fmt.Println(msg.Text)
	}
}
