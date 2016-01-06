package controllers

import (
	"ExamSystem/app/models"
	"log"
	"os"
	"strings"

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
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	var singleChoiceSummary map[string]int
	singleChoiceSummary, err = manager.GetSingleChoiceSummary()
	if err != nil {
		return this.RenderError(err)
	}

	var multipleChoiceSummary map[string]int
	multipleChoiceSummary, err = manager.GetMultipleChoiceSummary()
	if err != nil {
		return this.RenderError(err)
	}

	var trueFalseSummary map[string]int
	trueFalseSummary, err = manager.GetTrueFalseSummary()
	if err != nil {
		return this.RenderError(err)
	}

	trafficQuestionInfo := make(map[string]int)
	lawQuestionInfo := make(map[string]int)
	constructionQuestionInfo := make(map[string]int)

	if v, ok := singleChoiceSummary["交通"]; ok {
		trafficQuestionInfo["单选"] = v
	} else {
		trafficQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["交通"]; ok {
		trafficQuestionInfo["多选"] = v
	} else {
		trafficQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["交通"]; ok {
		trafficQuestionInfo["判断"] = v
	} else {
		trafficQuestionInfo["判断"] = 0
	}

	if v, ok := singleChoiceSummary["司法"]; ok {
		lawQuestionInfo["单选"] = v
	} else {
		lawQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["司法"]; ok {
		lawQuestionInfo["多选"] = v
	} else {
		lawQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["司法"]; ok {
		lawQuestionInfo["判断"] = v
	} else {
		lawQuestionInfo["判断"] = 0
	}

	if v, ok := singleChoiceSummary["建筑"]; ok {
		constructionQuestionInfo["单选"] = v
	} else {
		constructionQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["建筑"]; ok {
		constructionQuestionInfo["多选"] = v
	} else {
		constructionQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["建筑"]; ok {
		constructionQuestionInfo["判断"] = v
	} else {
		constructionQuestionInfo["判断"] = 0
	}

	results := make(map[string]map[string]int)
	results["交通"] = trafficQuestionInfo
	results["司法"] = lawQuestionInfo
	results["建筑"] = constructionQuestionInfo

	this.RenderArgs["results"] = results
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this Question) PostSingleChoice(singleChoice *models.SingleChoice) revel.Result {
	singleChoice.Type = strings.TrimSpace(singleChoice.Type)
	singleChoice.Discription = strings.TrimSpace(singleChoice.Discription)
	singleChoice.A = strings.TrimSpace(singleChoice.A)
	singleChoice.B = strings.TrimSpace(singleChoice.B)
	singleChoice.C = strings.TrimSpace(singleChoice.C)
	singleChoice.D = strings.TrimSpace(singleChoice.D)
	singleChoice.Answer = strings.TrimSpace(singleChoice.Answer)

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
	multipleChoice.Type = strings.TrimSpace(multipleChoice.Type)
	multipleChoice.Discription = strings.TrimSpace(multipleChoice.Discription)
	multipleChoice.A = strings.TrimSpace(multipleChoice.A)
	multipleChoice.B = strings.TrimSpace(multipleChoice.B)
	multipleChoice.C = strings.TrimSpace(multipleChoice.C)
	multipleChoice.D = strings.TrimSpace(multipleChoice.D)
	multipleChoice.E = strings.TrimSpace(multipleChoice.E)
	multipleChoice.F = strings.TrimSpace(multipleChoice.F)
	//multipleChoice.Answer = strings.TrimSpace(multipleChoice.Answer)

	this.Validation.Required(multipleChoice.Type).Message("请选择试题类别")
	this.Validation.Required(multipleChoice.Discription).Message("题目描述不能为空")
	this.Validation.Required(multipleChoice.A).Message("选项A不能为空")
	this.Validation.Required(multipleChoice.B).Message("选项B不能为空")
	this.Validation.Required(multipleChoice.C).Message("选项C不能为空")
	this.Validation.Required(multipleChoice.D).Message("选项D不能为空")
	this.Validation.Required(multipleChoice.E).Message("选项E不能为空")
	this.Validation.Required(multipleChoice.F).Message("选项F不能为空")
	//this.Validation.Required(len(answers)>0).Message("答案不能为空")
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
	trueFalse.Type = strings.TrimSpace(trueFalse.Type)
	trueFalse.Discription = strings.TrimSpace(trueFalse.Discription)
	trueFalse.Answer = strings.TrimSpace(trueFalse.Answer)

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

func (this Question) PostBatchSingleChoice(BatchSingleChoiceFile *os.File) revel.Result {
	return this.Render()
}

func (this Question) PostBatchMultipleChoice(BatchMultipleChoiceFile *os.File) revel.Result {
	return this.Render()
}

func (this Question) PostBatchTrueFalse(BatchTrueFalseFile *os.File) revel.Result {
	return this.Render()
}
