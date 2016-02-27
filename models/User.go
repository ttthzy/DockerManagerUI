package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserInfo struct {
	Id         bson.ObjectId `bson:"_id"`
	Account    string        `bson:"Account"`    //
	NickName   string        `bson:"NickName"`   //
	Phone      string        `bson:"Phone"`      //
	DockerUrl  []UserDocker  `bson:"DockerUrl"`  //
	IsDel      bool          `bson:"IsDel"`      //
	CreateDate string        `bson:"CreateDate"` //
}

/**
 * 添加UserInfo对象
 */
func AddUserInfo(p UserInfo) string {
	p.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := WitchCollection("UserInfo", query)
	if err != nil {
		return "false"
	}
	return p.Id.Hex()
}

/**
 * 更新记录
 */
func UpdateUserInfo(query bson.M, change bson.M) bool {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := WitchCollection("UserInfo", exop)
	if err != nil {
		return false
	}
	return true
}

/**
 * 获取所有记录
 */
func PageUserInfo() []UserInfo {
	var list []UserInfo
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&list)
	}
	err := WitchCollection("UserInfo", query)
	if err != nil {
		return list
	}
	return list
}

/**
 * 获取一条记录通过objectid
 */
func GetUserInfoById(id string) *UserInfo {
	objid := bson.ObjectIdHex(id)
	item := new(UserInfo)
	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&item)
	}
	WitchCollection("UserInfo", query)
	return item
}

/**
 * 获取一条记录通过phone
 */
func GetUserInfoByPhone(phone string) *UserInfo {
	item := new(UserInfo)
	query := func(c *mgo.Collection) error {
		qstr := bson.M{"phone": phone}
		return c.Find(qstr).One(&item)
	}
	WitchCollection("UserInfo", query)
	return item
}
