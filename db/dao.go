package db

import (
	"time"

	"github.com/flourish-ship/work-account/conf"
	mgo "gopkg.in/mgo.v2"
)

// DAOMongo ...
type DAOMongo struct {
	Session *mgo.Session
	db      *mgo.Database
}

type Result struct {
	Status
	Data interface{}
}

type Status int

const (
	Succuess Status = iota
	NotFound
	ValidationError
	DBError
	UnknownError
)

// NewDAOMongo ...
func NewDAOMongo(c *conf.DBConfig) (*DAOMongo, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{c.Addr},
		Timeout:  60 * time.Second,
		Database: c.Database,
		Username: c.Username,
		Password: c.Password,
	})
	if err != nil {
		return nil, err
	}
	return &DAOMongo{Session: session, db: session.DB("")}, nil
}
