package controllers

import "github.com/george518/PPGo_ApiAdmin/models"

type LogController struct {
	BaseController
}

func (self *LogController) Logs()  {
	//self.Ctx.WriteString(self.userName)
	models.Add("hello, this is my log",self.userName)

	//self.Data["pageTitle"] = "日志控制"
	//self.ServeJSONP()
	//self.display()
}

