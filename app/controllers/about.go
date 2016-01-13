package controllers

import "github.com/revel/revel"

type About struct {
	*revel.Controller
}

func (this About) About() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}
