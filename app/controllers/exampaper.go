package controllers

import (
	"examsystem/app/models"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/revel/revel"
)

type ExamPaper struct {
	*revel.Controller
}

func (this ExamPaper) Create() revel.Result {
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) PostCreate(examPaper *models.ExamPaper) revel.Result {
	examPaper.Type = strings.TrimSpace(examPaper.Type)
	//examPaper.CreateMethod = strings.TrimSpace(examPaper.CreateMethod)
	examPaper.Title = strings.TrimSpace(examPaper.Title)
	examPaper.Discription = strings.TrimSpace(examPaper.Discription)

	this.Validation.Required(examPaper.Type).Message("请选择试卷类别")
	//this.Validation.Required(examPaper.CreateMethod).Message("请选择试卷生成方式")
	this.Validation.Required(examPaper.Title).Message("请填写试卷标题")
	this.Validation.Required(examPaper.Discription).Message("请填写试卷描述")
	this.Validation.Required(examPaper.Score > 0).Message("请设置试卷总分数(大于零)")
	this.Validation.Required(examPaper.Time > 0).Message("请设置考试时间(大于零)")
	this.Validation.Required(examPaper.SCCount > 0).Message("请设置单选题数量(大于零)")
	this.Validation.Required(examPaper.SCScore > 0).Message("请设置单选题每题分值(大于零)")
	this.Validation.Required(examPaper.MCCount > 0).Message("请设置多选题数量(大于零)")
	this.Validation.Required(examPaper.MCScore > 0).Message("请设置多选题每题分值(大于零)")
	this.Validation.Required(examPaper.TFCount > 0).Message("请设置判断题数量(大于零)")
	this.Validation.Required(examPaper.TFScore > 0).Message("请设置判断题每题分值(大于零)")

	examPaper.CreateMethod = "随机生成"
	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(ExamPaper.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}
	defer manager.Close()

	sc := []models.SingleChoice{}
	sc, err = manager.GetRandomSingleChoice(examPaper.Type, examPaper.SCCount)
	log.Println(sc)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	mc := []models.MultipleChoice{}
	mc, err = manager.GetRandomMultipleChoice(examPaper.Type, examPaper.MCCount)
	log.Println(mc)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	tf := []models.TrueFalse{}
	tf, err = manager.GetRandomTrueFalse(examPaper.Type, examPaper.TFCount)
	log.Println(tf)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	examPaper.SC = sc
	examPaper.MC = mc
	examPaper.TF = tf

	t := time.Now()
	examPaper.TimeStamp = t.Format("2006-01-02 15:04:05")
	examPaper.IDCode = "HGY" + strconv.Itoa(t.Year()) + strconv.Itoa((int)(t.Month())) +
		strconv.Itoa(t.Day()) + strconv.Itoa(t.Hour()) +
		strconv.Itoa(t.Minute()) + strconv.Itoa(t.Second())

	err = manager.AddExamPaper(examPaper)
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	this.Flash.Success("试卷成功生成")

	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Redirect(ExamPaper.Create)
}

func (this ExamPaper) saveExamPaperFile(examPaperFile *os.File) (filePath string, err error) {
	//// 使用revel request formfile获取文件数据
	//file, handler, err := this.Request.FormFile("examPaperFile")
	//if err != nil {
	//	this.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//// 读取所有数据
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	this.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//
	//// 获取当前路径
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	this.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//
	//// 文件路径
	//filePath = dir + "/" + handler.Filename
	//
	//// 保存到文件
	//err = ioutil.WriteFile(filePath, data, 0777)
	//if err != nil {
	//	this.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//
	//return filePath, nil
	// TODO for debug
	return "nil", nil
}

func (this ExamPaper) PostUpload(file *os.File, pType string) revel.Result {
	filePath, err := this.saveExamPaperFile(file)
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	examPaper, scFilePath, mcFilePath, tfFilePath, err := models.ParseExamPaperFile(filePath)
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	scf, err := os.Open(scFilePath)
	if err != nil {
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}
	defer scf.Close()
	mcf, err := os.Open(mcFilePath)
	if err != nil {
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}
	defer mcf.Close()
	tff, err := os.Open(tfFilePath)
	if err != nil {
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}
	defer tff.Close()

	scs, err := models.ParseSingleChoiceFile(scf, pType)
	mcs, err := models.ParseMultipleChoiceFile(mcf, pType)
	tfs, err := models.ParseTrueFalseFile(tff, pType)

	t := time.Now()
	examPaper.Type = pType
	examPaper.TimeStamp = t.Format("2006-01-02 15:04:05")
	examPaper.IDCode = "HGY" + strconv.Itoa(t.Year()) + strconv.Itoa((int)(t.Month())) +
		strconv.Itoa(t.Day()) + strconv.Itoa(t.Hour()) +
		strconv.Itoa(t.Minute()) + strconv.Itoa(t.Second())
	examPaper.CreateMethod = "套题上传"
	examPaper.SCCount = len(scs)
	examPaper.SC = scs
	examPaper.MCCount = len(mcs)
	examPaper.MC = mcs
	examPaper.TFCount = len(tfs)
	examPaper.TF = tfs

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}
	defer manager.Close()

	err = manager.AddExamPaper(&examPaper)
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Create)
	}

	this.Flash.Success("试卷成功上传")

	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Redirect(ExamPaper.Create)
}

func (this ExamPaper) Preview(title string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}
	defer manager.Close()

	examPaper, e := manager.GetExamPaperByTitle(title)
	if e != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}

	this.ViewArgs["scCount"] = len(examPaper.SC)
	this.ViewArgs["mcCount"] = len(examPaper.MC)
	this.ViewArgs["tfCount"] = len(examPaper.TF)
	this.ViewArgs["examPaper"] = examPaper
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) View() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}
	defer manager.Close()

	examPapers, e := manager.GetAllExamPaper()

	if e != nil {
		log.Println(e)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}

	this.ViewArgs["examPapers"] = examPapers
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) Publish() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}
	defer manager.Close()

	examPapers, e := manager.GetAllExamPaper()

	if e != nil {
		log.Println(e)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}

	this.ViewArgs["examPapers"] = examPapers
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) PostPublish(exmpaperTitle string) revel.Result {
	exmpaperTitle = strings.TrimSpace(exmpaperTitle)
	this.Validation.Required(exmpaperTitle).Message("请选择一个试题")
	log.Println(exmpaperTitle)

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(ExamPaper.Publish)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Publish)
	}
	defer manager.Close()

	examinees, err := manager.GetAllExaminee()
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Publish)
	}

	examPaper, err := manager.GetExamPaperByTitle(exmpaperTitle)
	if err != nil {
		log.Println(err)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Redirect(ExamPaper.Publish)
	}

	for _, examinee := range examinees {
		examPaper.Status = models.UnFinished
		models.ChaosExamPaper(&examPaper)
		examinee.ExamPaper = examPaper

		err := manager.UpdateExaminee(&examinee)
		if err != nil {
			log.Println(err)
			this.Response.Status = 500
			this.Flash.Error(err.Error())
			return this.Redirect(ExamPaper.Publish)
		}
	}

	this.Flash.Success("考试成功发布")

	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Redirect(ExamPaper.Publish)
}

func (this ExamPaper) Score(idCard string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}
	defer manager.Close()

	examinee, e := manager.GetExamineeByIDCard(idCard)

	if e != nil {
		log.Println(e)
		this.Response.Status = 500
		this.Flash.Error(err.Error())
		return this.Render()
	}

	log.Println("查到考生信息: ", examinee)

	this.ViewArgs["examinee"] = examinee
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) QueryScore() revel.Result {
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this ExamPaper) PostQueryScore(examineeIDCard string) revel.Result {
	examineeIDCard = strings.TrimSpace(examineeIDCard)

	this.Validation.Required(examineeIDCard).Message("请输入身份证号码")
	this.Validation.Length(examineeIDCard, 18).Message("身份证号有误")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(ExamPaper.QueryScore)
	}

	return this.Redirect("/ExamPaper/Score/%s", examineeIDCard)
	//return this.Redirect(ExamPaper.Score("huguangyue"))
}
