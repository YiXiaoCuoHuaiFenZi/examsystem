package controllers

import (
	"ExamSystem/app/models"
	"log"
	//"strconv"

	"github.com/revel/revel"
)

type ExamPaper struct {
	*revel.Controller
}

func (this ExamPaper) Create() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) PostCreate(examPaper *models.ExamPaper) revel.Result {
	this.Validation.Required(examPaper.Type).Message("请选择试卷类别")
	this.Validation.Required(examPaper.CreateMethod).Message("请选择试卷生成方式")
	this.Validation.Required(examPaper.Title).Message("请填写试卷标题")
	this.Validation.Required(examPaper.Discription).Message("请填写试卷描述")
	this.Validation.Required(examPaper.SCCount).Message("请设置单选题数量")
	this.Validation.Required(examPaper.SCScore).Message("请设置单选题每题分值")
	this.Validation.Required(examPaper.MCCount).Message("请设置多选题数量")
	this.Validation.Required(examPaper.MCScore).Message("请设置多选题每题分值")
	this.Validation.Required(examPaper.TFCount).Message("请设置判断题数量")
	this.Validation.Required(examPaper.TFScore).Message("请设置判断题每题分值")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(ExamPaper.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	sc := []models.SingleChoice{}
	sc, err = manager.GetRandomSingleChoice(examPaper.SCCount)
	log.Println(sc)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		return this.RenderError(err)
	}

	mc := []models.MultipleChoice{}
	mc, err = manager.GetRandomMultipleChoice(examPaper.MCCount)
	log.Println(mc)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		return this.RenderError(err)
	}

	tf := []models.TrueFalse{}
	tf, err = manager.GetRandomTrueFalse(examPaper.TFCount)
	log.Println(tf)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		return this.RenderError(err)
	}

	examPaper.SC = sc
	examPaper.MC = mc
	examPaper.TF = tf

	err = manager.AddExamPaper(examPaper)
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}

	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Redirect(ExamPaper.Create)
}

func (this ExamPaper) View() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) Publish() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) Score(idCard string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	examinee, e := manager.GetExamineeByIDCard(idCard)

	if e != nil {
		log.Println(e)
		this.Response.Status = 500
		return this.RenderError(e)
	}

	log.Println("查到考生信息: ", examinee)

	this.RenderArgs["examinee"] = examinee
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) QueryScore() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) PostQueryScore(examineeIDCard string) revel.Result {
	this.Validation.Required(examineeIDCard).Message("请输入身份证号码")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(ExamPaper.QueryScore)
	}

	return this.Redirect("/ExamPaper/Score/%s", examineeIDCard)
	//return this.Redirect(ExamPaper.Score("huguangyue"))
}
