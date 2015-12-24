package controllers

import (
//	"ExamSystem/app/models"
//	"log"

	"github.com/revel/revel"
)

type Examinee struct {
	*revel.Controller
}

func (this Examinee) Info() revel.Result {
	return this.Render()
}
