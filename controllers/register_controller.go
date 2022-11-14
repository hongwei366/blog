package controllers

import (
	"blog/models"
	"blog/utils"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
)

// RegisterController /*注册控制器*/
type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	//获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	log.Info("请求参数username=", username, "password=", password)
	//判断用户是否注册过
	id := models.QueryUserWithParam(username, password)
	log.Info("id={}", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 201, "message": "用户已注册过，请直接登录"}
		c.ServeJSON()
		return
	}
	//写入注册用户
	password = utils.Md5(password)
	_, err := models.InsertUser(models.User{Username: username, Password: password, Status: 0, Createtime: time.Now().Unix()})
	if err != nil {
		log.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 200, "message": "注册成功"}
	}
	c.ServeJSON()
}
