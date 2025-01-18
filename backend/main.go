package main

import (
	"financify/bulk-entry/consts"
	"fmt"
	"path"
)

func main() {
	conf := &consts.Configuration{}
	conf.LoadConfig(path.Dir("config"))
	fmt.Printf("%+v", conf)
}
