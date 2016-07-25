package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id         bson.ObjectId `json:"id"`
	UserName   string        `json:"userName"`   //姓名
	Gender     string        `json:"gender"`     //性别
	NickName   string        `json:"nickName"`   //外号
	PinYin     string        `json:"pinyin"`     //简拼
	FullPinYin string        `json:"fullPinYin"` //全拼
	Password   string        `json:"passWord"`
}
