package controllers

import (
//	"ExamSystem/app/models"
//	"log"

	"github.com/revel/revel"
)

type ExamPaper struct {
	*revel.Controller
}

func (this ExamPaper) Create() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}

func (this ExamPaper) View() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}

func (this ExamPaper) Publish() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}

func (this ExamPaper) Score() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}
