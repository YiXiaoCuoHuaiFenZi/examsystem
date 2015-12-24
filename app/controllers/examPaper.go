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
