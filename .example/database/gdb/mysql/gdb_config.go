package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	if r, err := g.DB().Model("user").Where("uid=?", 1).One(); err == nil {
		fmt.Println(r["uid"].Int())
		fmt.Println(r["name"].String())
	} else {
		fmt.Println(err)
	}
}
