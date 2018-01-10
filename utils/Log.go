package utils

import "github.com/astaxie/beego/logs"

var LOG *logs.BeeLogger

func init()  {
	//初始化 日志
	LOG = logs.NewLogger()
	LOG.SetLogger(logs.AdapterConsole)
	LOG.EnableFuncCallDepth(true)
}
