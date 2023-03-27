package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (this App) Index() revel.Result {
	return this.Render()
}

func (this App) Introduce() revel.Result {
	return this.Render()
}

func (this App) About() revel.Result {
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}
