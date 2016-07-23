package main

import (
	"github.com/flourish-ship/work-account/conf"
	"github.com/flourish-ship/work-account/daomongo"
	"github.com/flourish-ship/work-account/routers"
)

func main() {
	c := getConf("./config.json")

	dao := getDAOMongo(c.DB)
	defer dao.Session.Close()

	routers.NewAccountManager(dao, c.API).Server()
}

func getConf(path string) *conf.Config {
	c := &conf.Config{}
	err := conf.LoadConfig(path, c)
	if err != nil {
		panic(err.Error())
	}
	return c
}

func getDAOMongo(c *conf.DBConfig) *daomongo.DAOMongo {
	dao, err := daomongo.NewDAOMongo(c)
	if err != nil {
		panic(err.Error())
	}
	return dao
}
