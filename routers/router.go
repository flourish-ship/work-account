package routers

import (
	"github.com/flourish-ship/work-account/conf"
	"github.com/flourish-ship/work-account/idao"
	"github.com/iris-contrib/sessiondb/redis"
	"github.com/iris-contrib/sessiondb/redis/service"
	"github.com/kataras/iris"
)

// AccountManager ...
type AccountManager struct {
	Redis *redis.Database
	API   *iris.Framework
	DAO   idao.IDAO
}

// NewAccountManager ...
func NewAccountManager(daoImpl idao.IDAO, c *conf.Config) *AccountManager {
	//iris.UseSessionDB(db)
	return &AccountManager{
		Redis: initReils(c),
		API:   iris.New(),
		DAO:   daoImpl,
	}
}

func initReils(c *conf.Config) *redis.Database {
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
	am.initRouter()
}

func (am *AccountManager) initRouter() {
	api := am.API
	api.Party("/v1/am")
	{
		api.Post("/login", am.Login)
	}
}

// Server ...
func (am *AccountManager) Server() {
	am.API.Listen(":3030")
}
