package controllers

import (
//	"ExamSystem/app/models"
//	"log"

	"github.com/revel/revel"
)

type Question struct {
	*revel.Controller
}

func (this Question) Create() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}

func (this Question) View() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}
