package controllers

import "github.com/revel/revel"

type About struct {
	*revel.Controller
}

func (this About) About() revel.Result {
	return this.Render()
}
