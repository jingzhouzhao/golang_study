package tailf

import (
	"time"
	"github.com/astaxie/beego/logs"
	"fmt"
	"errors"
	"github.com/hpcloud/tail"
)

type TextMsg struct{
	Msg string
	Topic string
}

type CollectConfig struct{
	LogPath string
	Topic string
	ChanSize int
}

type TailObj struct{
	tail *tail.Tail
	collectConfig CollectConfig
	msgs chan *TextMsg
}

type TailObjMgr struct{
	tailObjs [] *TailObj
}

var (
	tailObjMgr TailObjMgr
)

func GetMsg() (msg *TextMsg){
	for i := 0; i < len(tailObjMgr.tailObjs); i++ {
		select {
		case msg=<-tailObjMgr.tailObjs[i].msgs:
			logs.Trace("get msg from %s chan", tailObjMgr.tailObjs[i].collectConfig.LogPath)
		}
	}
	return 
}

func InitTail(collectConfigs []CollectConfig) (err error){
	if len(collectConfigs)==0{
		err = errors.New("collectConfigs must not be empty")
		return
	}
	for _,v:=range collectConfigs{
		var tailObj TailObj
		tailObj.collectConfig = v
		tails,tailErr:=tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			//Location:&tail.SeekInfo{Offset:0,Whence:2},
			MustExist: false,
			Poll:      true,
		})
		if tailErr !=nil {
			fmt.Println("tail file failed:",tailErr)
			err = tailErr
			return
		}
		tailObj.tail = tails
		tailObj.msgs = make(chan *TextMsg,v.ChanSize)
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, &tailObj)
		go readFromTail(tailObj)
	}
	logs.Debug("initialize tailf Success")
	return
}

func readFromTail(tailObj TailObj){
	for{
		line,ok := <-tailObj.tail.Lines
		if !ok{
			logs.Warn("tail file close reopen,filename:", tailObj.tail.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		textMsg :=&TextMsg{
			Msg:line.Text,
			Topic:tailObj.collectConfig.Topic,
		}
		tailObj.msgs <-textMsg
	}
}