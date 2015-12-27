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
	return this.Render()
}

func (this ExamPaper) View() revel.Result {
	return this.Render()
}

func (this ExamPaper) Publish() revel.Result {
	return this.Render()
}

func (this ExamPaper) Score() revel.Result {
	return this.Render()
}
