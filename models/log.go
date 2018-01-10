package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/george518/PPGo_ApiAdmin/utils"
	"time"
)

type Log struct {
	Id      int
	Time    string
	Message string
	User    string
}


func (a *Log) TableName() string {
	return TableName("log")
}

func  Add(message ,user string)  {
	t:=time.Now()
	l:=Log{Message:message,User:user,Time:t.String()}

	utils.LOG.Debug("",l)
	index,e:=orm.NewOrm().Insert(l)

	if e!=nil {
		panic(e)
	}

	utils.LOG.Debug("message :",index)
}



