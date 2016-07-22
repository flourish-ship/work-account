package main

import (
	"github.com/flourish-ship/work-account/conf"
	"github.com/flourish-ship/work-account/daomongo"
	"github.com/flourish-ship/work-account/routers"
)

func main() {
	c := getConf("./config.json")
	routers.NewAccountManager(daomongo.NewDAOMongo(c), c).Server()
}

func getConf(path string) *conf.Config {
	c := &conf.Config{}
	err := conf.LoadConfig(path, c)
	if err != nil {
		panic(err.Error())
	}
	return c
}
