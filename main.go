// @APIVersion 1.0.0
// @APITitle Account management of work attendance
// @APIDescription It's a part of work attendance project, and development is in-progress.
// @License Apache Licence Vesion 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0
// @BasePath /am/v1/
// @SubApi Account related(eg. signin,signout) API [/account]
package main

import (
	"github.com/flourish-ship/work-account/conf"
	"github.com/flourish-ship/work-account/db"
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

func getDAOMongo(c *conf.DBConfig) *db.DAOMongo {
	dao, err := db.NewDAOMongo(c)
	if err != nil {
		panic(err.Error())
	}
	return dao
}
