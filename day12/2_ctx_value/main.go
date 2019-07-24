package main

import (
	"fmt"
	"context"
)

func process(ctx context.Context){
	ret,ok := ctx.Value("trace_id").(int)
	if !ok{
		fmt.Println("invalid trace_id!")
		return
	}
	fmt.Println("trace_id:",ret)
	sid,ok:=ctx.Value("session_id").(string)
	if !ok{
		fmt.Println("invalid session_id!")
		return
	}
	fmt.Println("session_id:",sid)

}

func main(){
	ctx:=context.WithValue(context.Background(), "trace_id", 123456)
	//process(ctx)
	//ctx继承
	ctx=context.WithValue(ctx, "session_id", "jwjeroiwdlknflasjdfoiwerljkasdjfa")
	process(ctx)
}