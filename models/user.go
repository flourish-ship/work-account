package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username     string        `bson:"username" json:"username"`
	PasswordHash string        `bson:"password_hash" json:"-"`
	Password     string        `bson:"-" json:"password"`
	Gender       string        `bson:"gender,omitempty" json:"gender"`
	NickName     string        `bson:"nick_name,omitempty" json:"nickName"`
	PinYin       string        `json:"pinyin,omitempty" json:"pinyin"`
}
