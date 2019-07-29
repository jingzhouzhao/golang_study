package tailf

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type TextMsg struct {
	Msg   string
	Topic string
}

type CollectConfig struct {
	LogPath  string
	Topic    string
	ChanSize int
}

type TailObj struct {
	tail          *tail.Tail
	collectConfig CollectConfig
	msgs          chan *TextMsg
	ctx           context.Context
	cancle func()
}

type TailObjMgr struct {
	tailObjs []*TailObj
}

var (
	tailObjMgr TailObjMgr
)

func GetMsg() (msg *TextMsg) {
	for i := 0; i < len(tailObjMgr.tailObjs); i++ {
		select {
		case msg = <-tailObjMgr.tailObjs[i].msgs:
			//logs.Debug("get msg from %s chan", tailObjMgr.tailObjs[i].collectConfig.LogPath)
		default:
			//logs.Debug("no msg got from %s chan", tailObjMgr.tailObjs[i].collectConfig.LogPath)
		}
	}
	return
}

func InitTail(collectConfigs []CollectConfig) (err error) {
	if len(collectConfigs) == 0 {
		//err = errors.New("collectConfigs must not be empty")
		logs.Warn("collectConfigs is empty!")
		return
	}
	for _, v := range collectConfigs {
		logs.Debug("new tial file:%#v",v)
		newTailFile(v)
	}
	logs.Debug("initialize tailf Success")
	return
}

func newTailFile(v CollectConfig) (to TailObj,err error){
	var tailObj  = TailObj{
		collectConfig:v,
	}
	ctx,cancle:=context.WithCancel(context.Background())
	tailObj.ctx = ctx
	tailObj.cancle = cancle
	tails, err := tail.TailFile(v.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Error("tail file failed:", err)
		return
	}
	tailObj.tail = tails
	tailObj.msgs = make(chan *TextMsg, v.ChanSize)
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, &tailObj)
	go readFromTail(tailObj)
	to = tailObj
	return
}

func readFromTail(tailObj TailObj) {
	for {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen,filename:", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.collectConfig.Topic,
		}
		select {
		case <-tailObj.ctx.Done():
			//cancle
			return
		case tailObj.msgs <- textMsg:
		}

	}
}

func UpdateCollectConfigs(collectConfs []CollectConfig) {
	var tailObjs []*TailObj
	for _,t:=range tailObjMgr.tailObjs{
		exists:=false
		for _,c:=range collectConfs{
			if c.LogPath == t.collectConfig.LogPath{
				exists = true
				//更新配置
				t.collectConfig = c
				tailObjs = append(tailObjs, t)
				break
			}
		}
		if !exists {
			//删除配置
			t.cancle()
		}
	}
	//新增配置
	if len(tailObjMgr.tailObjs)<=len(collectConfs){
		for _,c:=range collectConfs{
			exists:=false
			for _,t:=range tailObjMgr.tailObjs{
				if c.LogPath == t.collectConfig.LogPath{
					exists = true
					break
				}
			}
			if !exists{
				tailObj,err:=newTailFile(c)
				if err!=nil{
					logs.Error("new tail file failed:", err)
					continue
				}
				tailObjs = append(tailObjs, &tailObj)
			}
		}
	}
	logs.Debug("after update collect configs:%#v", tailObjs)
	tailObjMgr.tailObjs = tailObjs
}
