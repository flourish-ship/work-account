package routers

import (
	"github.com/flourish-ship/work-account/conf"
	"github.com/flourish-ship/work-account/idao"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/sessiondb/redis"
	"github.com/iris-contrib/sessiondb/redis/service"
	"github.com/kataras/iris"
)

const (
	// PERFIX ...
	PERFIX = "/am/v1"
)

// AccountManager ...
type AccountManager struct {
	config *conf.APIConfig
	Redis  *redis.Database
	API    *iris.Framework
	DAO    idao.IDAO
}

// Register ...
type Register interface {
	Registe(*AccountManager)
}

// NewAccountManager ...
func NewAccountManager(daoImpl idao.IDAO, c *conf.APIConfig) *AccountManager {
	return &AccountManager{
		config: c,
		Redis:  initReils(c),
		API:    iris.New(),
		DAO:    daoImpl,
	}
}

func initReils(c *conf.APIConfig) *redis.Database {
	return redis.New(service.Config{
		Network:       service.DefaultRedisNetwork,
		Addr:          c.Redis.Addr,
		Password:      c.Redis.Password,
		Database:      c.Redis.Database,
		MaxIdle:       0,
		MaxActive:     0,
		IdleTimeout:   service.DefaultRedisIdleTimeout,
		Prefix:        "",
		MaxAgeSeconds: service.DefaultRedisMaxAgeSeconds,
	})
}

func (am *AccountManager) initialize() {
	am.API.UseSessionDB(am.Redis)
	am.API.Use(logger.New(iris.Logger))
	am.API.StaticWeb("/docs", "./swagger/", 1)
	am.rigiste(&AccountRouter{R: "/account"})
}

func (am *AccountManager) rigiste(registers ...Register) {
	if registers == nil || len(registers) == 0 {
		return
	}
	for _, register := range registers {
		register.Registe(am)
	}
}

// Server ...
func (am *AccountManager) Server() {
	am.initialize()
	am.API.Listen(am.config.Port)
}
