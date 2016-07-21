package main

import (
	"github.com/flourish-ship/work-account/daomongo"
	"github.com/flourish-ship/work-account/routers"
)

func main() {
	routers.NewAccountManager(daomongo.NewDAOMongo()).Server()
}
