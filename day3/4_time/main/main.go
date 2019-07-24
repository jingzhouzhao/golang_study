package main

import (
	"fmt"
	"time"
)

func main()  {
	now:=time.Now()
	nowstr:=now.Format("2006/01/02 15:04:05")
	fmt.Printf("now:%s\n",nowstr)

	start:=time.Now().UnixNano()

	time.Sleep(2*time.Second)

	end:=time.Now().UnixNano()
	fmt.Printf("花费了%vs\n",(end-start)/int64(time.Second))
	fmt.Printf("一微妙等于%d纳秒\n",int64(time.Microsecond))
	fmt.Printf("一毫秒等于%d纳秒\n",int64(time.Millisecond))
	var cost int64  = 10
	fmt.Printf("转换int64为秒时间,等于%s\n", time.Duration(cost)*time.Second)
	fmt.Printf("转换int64为毫秒时间,等于%s\n", time.Duration(cost)*time.Millisecond)


}