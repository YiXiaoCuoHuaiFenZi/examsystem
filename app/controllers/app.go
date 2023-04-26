package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (ap App) Index() revel.Result {
	return ap.Render()
}

func (ap App) Introduce() revel.Result {
	return ap.Render()
}

func (ap App) About() revel.Result {
	ap.ViewArgs["adminIDCard"] = ap.Session["adminIDCard"]
	ap.ViewArgs["adminName"] = ap.Session["adminName"]

	return ap.Render()
}
