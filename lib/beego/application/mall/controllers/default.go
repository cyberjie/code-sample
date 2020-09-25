package controllers

import "github.com/astaxie/beego"

type DefaultController struct {
	beego.Controller
}
func (this *DefaultController)Index(){
	this.Data["code"]=200
	this.TplName="index.html"
}