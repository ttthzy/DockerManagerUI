package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const URL = "mongodb://admin:H2Xv6cznmCm2@mongodb-ttthzygi35.tenxcloud.net:44329" //mongodb连接字符串(外网)
//const URL = "mongodb://admin:H2Xv6cznmCm2@172.16.35.78:27017" //mongodb连接字符串(内网)

var (
	mgoSession *mgo.Session
	dataBase   = "docker"
	FmartDate  = "2006-01-02 15:04"
)

///将当前时间转化为字符串
func TimeNowToStr() string {
	return time.Now().Format(FmartDate)
}

///将UTC时间增加8小时为东部时间
func TimeUtcToCst(t time.Time) time.Time {
	return t.Add(time.Hour * time.Duration(8))
}

///将当前UTC时间增加8小时为东部时间
func TimeNowUtcToCst() time.Time {
	return time.Now().Add(time.Hour * time.Duration(8))
}

///将str转换为时间格式
func StrToTime(st string) time.Time {
	t, _ := time.ParseInLocation(FmartDate, st, time.Local)
	return t
}

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

//公共方法，获取collection对象
func WitchCollection(collection string, s func(*mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func SearchDB(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = WitchCollection(collectionName, exop)
	return
}
