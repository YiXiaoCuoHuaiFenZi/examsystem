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
	return this.Render()
}

func (this Question) View() revel.Result {
	return this.Render()
}
