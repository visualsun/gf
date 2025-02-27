package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gpool"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	// 创建对象复用池，对象过期时间为3000毫秒，并给定创建及销毁方法
	p := gpool.New(3000*time.Millisecond, func() (interface{}, error) {
		return gtcp.NewConn("www.baidu.com:80")
	}, func(i interface{}) {
		glog.Print("expired")
		i.(*gtcp.Conn).Close()
	})
	conn, err := p.Get()
	if err != nil {
		panic(err)
	}
	result, err := conn.(*gtcp.Conn).SendRecv([]byte("HEAD / HTTP/1.1\n\n"), -1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	// 丢回池中以便重复使用
	p.Put(conn)
	// 等待一定时间观察过期方法调用
	time.Sleep(4 * time.Second)
}
