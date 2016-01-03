package controllers

import (
	"ExamSystem/app/models"
	"log"

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

func (this Question) PostSingleChoice(singleChoice *models.SingleChoice) revel.Result {
	this.Validation.Required(singleChoice.Type).Message("请选择试题类别")
	this.Validation.Required(singleChoice.Discription).Message("题目描述不能为空")
	this.Validation.Required(singleChoice.A).Message("选项A不能为空")
	this.Validation.Required(singleChoice.B).Message("选项B不能为空")
	this.Validation.Required(singleChoice.C).Message("选项C不能为空")
	this.Validation.Required(singleChoice.D).Message("选项D不能为空")
	this.Validation.Required(singleChoice.Answer).Message("答案不能为空")

	switch singleChoice.Answer {
	case "A":
		singleChoice.Answer = singleChoice.A
	case "B":
		singleChoice.Answer = singleChoice.B
	case "C":
		singleChoice.Answer = singleChoice.C
	case "D":
		singleChoice.Answer = singleChoice.D
	default:
		break
	}

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.AddSingleChoice(singleChoice)
	if err != nil {
		this.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "singleChoice.Discription"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	this.Session["createStatus"] = "success"
	log.Println("创建题目成功：", singleChoice)

	return this.Redirect(Question.Create)
}

func (this Question) PostMultipleChoice(multipleChoice *models.MultipleChoice, answers []string) revel.Result {
	this.Validation.Required(multipleChoice.Type).Message("请选择试题类别")
	this.Validation.Required(multipleChoice.Discription).Message("题目描述不能为空")
	this.Validation.Required(multipleChoice.A).Message("选项A不能为空")
	this.Validation.Required(multipleChoice.B).Message("选项B不能为空")
	this.Validation.Required(multipleChoice.C).Message("选项C不能为空")
	this.Validation.Required(multipleChoice.D).Message("选项D不能为空")
	this.Validation.Required(multipleChoice.E).Message("选项E不能为空")
	this.Validation.Required(multipleChoice.F).Message("选项F不能为空")
	this.Validation.Required(answers).Message("答案不能为空")

	var as []string
	for _, a := range answers {
		switch a {
		case "A":
			as = append(as, multipleChoice.A)
		case "B":
			as = append(as, multipleChoice.B)
		case "C":
			as = append(as, multipleChoice.C)
		case "D":
			as = append(as, multipleChoice.D)
		case "E":
			as = append(as, multipleChoice.E)
		case "F":
			as = append(as, multipleChoice.F)
		default:
			break
		}
	}
	multipleChoice.Answer = as

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.AddMultipleChoice(multipleChoice)
	if err != nil {
		this.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "multipleChoice.Discription"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	this.Session["createStatus"] = "success"
	log.Println("创建题目成功：", multipleChoice)

	return this.Redirect(Question.Create)
}

func (this Question) PostTrueFalse(trueFalse *models.TrueFalse) revel.Result {
	this.Validation.Required(trueFalse.Type).Message("请选择试题类别")
	this.Validation.Required(trueFalse.Discription).Message("题目描述不能为空")
	this.Validation.Required(trueFalse.Answer).Message("答案不能为空")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.AddTrueFalse(trueFalse)
	if err != nil {
		this.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "trueFalse.Discription"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Question.Create)
	}

	this.Session["createStatus"] = "success"
	log.Println("创建题目成功：", trueFalse)

	return this.Redirect(Question.Create)
}
