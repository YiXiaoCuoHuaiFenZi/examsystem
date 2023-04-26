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

func (ep ExamPaper) Create() revel.Result {
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) PostCreate(examPaper *models.ExamPaper) revel.Result {
	examPaper.Type = strings.TrimSpace(examPaper.Type)
	//examPaper.CreateMethod = strings.TrimSpace(examPaper.CreateMethod)
	examPaper.Title = strings.TrimSpace(examPaper.Title)
	examPaper.Description = strings.TrimSpace(examPaper.Description)

	ep.Validation.Required(examPaper.Type).Message("请选择试卷类别")
	//ep.Validation.Required(examPaper.CreateMethod).Message("请选择试卷生成方式")
	ep.Validation.Required(examPaper.Title).Message("请填写试卷标题")
	ep.Validation.Required(examPaper.Description).Message("请填写试卷描述")
	ep.Validation.Required(examPaper.Score > 0).Message("请设置试卷总分数(大于零)")
	ep.Validation.Required(examPaper.Time > 0).Message("请设置考试时间(大于零)")
	ep.Validation.Required(examPaper.SCCount > 0).Message("请设置单选题数量(大于零)")
	ep.Validation.Required(examPaper.SCScore > 0).Message("请设置单选题每题分值(大于零)")
	ep.Validation.Required(examPaper.MCCount > 0).Message("请设置多选题数量(大于零)")
	ep.Validation.Required(examPaper.MCScore > 0).Message("请设置多选题每题分值(大于零)")
	ep.Validation.Required(examPaper.TFCount > 0).Message("请设置判断题数量(大于零)")
	ep.Validation.Required(examPaper.TFScore > 0).Message("请设置判断题每题分值(大于零)")

	examPaper.CreateMethod = "随机生成"
	if ep.Validation.HasErrors() {
		ep.Validation.Keep()
		ep.FlashParams()
		return ep.Redirect(ExamPaper.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}
	defer manager.Close()

	sc := []models.SingleChoice{}
	sc, err = manager.GetRandomSingleChoice(examPaper.Type, examPaper.SCCount)
	log.Println(sc)
	if err != nil {
		log.Println(err)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	mc := []models.MultipleChoice{}
	mc, err = manager.GetRandomMultipleChoice(examPaper.Type, examPaper.MCCount)
	log.Println(mc)
	if err != nil {
		log.Println(err)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	tf := []models.TrueFalse{}
	tf, err = manager.GetRandomTrueFalse(examPaper.Type, examPaper.TFCount)
	log.Println(tf)
	if err != nil {
		log.Println(err)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
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
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	ep.Flash.Success("试卷成功生成")

	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Redirect(ExamPaper.Create)
}

func (ep ExamPaper) saveExamPaperFile(examPaperFile *os.File) (filePath string, err error) {
	//// 使用revel request formfile获取文件数据
	//file, handler, err := ep.Request.FormFile("examPaperFile")
	//if err != nil {
	//	ep.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//// 读取所有数据
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	ep.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//
	//// 获取当前路径
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	ep.Response.Status = 500
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
	//	ep.Response.Status = 500
	//	log.Println(err)
	//	return "", err
	//}
	//
	//return filePath, nil
	// TODO for debug
	return "nil", nil
}

func (ep ExamPaper) PostUpload(file *os.File, pType string) revel.Result {
	filePath, err := ep.saveExamPaperFile(file)
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	examPaper, scFilePath, mcFilePath, tfFilePath, err := models.ParseExamPaperFile(filePath)
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	scf, err := os.Open(scFilePath)
	if err != nil {
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}
	defer scf.Close()
	mcf, err := os.Open(mcFilePath)
	if err != nil {
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}
	defer mcf.Close()
	tff, err := os.Open(tfFilePath)
	if err != nil {
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
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
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}
	defer manager.Close()

	err = manager.AddExamPaper(&examPaper)
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Create)
	}

	ep.Flash.Success("试卷成功上传")

	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Redirect(ExamPaper.Create)
}

func (ep ExamPaper) Preview(title string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}
	defer manager.Close()

	examPaper, e := manager.GetExamPaperByTitle(title)
	if e != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}

	ep.ViewArgs["scCount"] = len(examPaper.SC)
	ep.ViewArgs["mcCount"] = len(examPaper.MC)
	ep.ViewArgs["tfCount"] = len(examPaper.TF)
	ep.ViewArgs["examPaper"] = examPaper
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) View() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}
	defer manager.Close()

	examPapers, e := manager.GetAllExamPaper()

	if e != nil {
		log.Println(e)
		ep.Response.Status = 500
		ep.Flash.Error(e.Error())
		return ep.Render()
	}

	ep.ViewArgs["examPapers"] = examPapers
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) Publish() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}
	defer manager.Close()

	examPapers, e := manager.GetAllExamPaper()

	if e != nil {
		log.Println(e)
		ep.Response.Status = 500
		ep.Flash.Error(e.Error())
		return ep.Render()
	}

	ep.ViewArgs["examPapers"] = examPapers
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) PostPublish(exmpaperTitle string) revel.Result {
	exmpaperTitle = strings.TrimSpace(exmpaperTitle)
	ep.Validation.Required(exmpaperTitle).Message("请选择一个试题")
	log.Println(exmpaperTitle)

	if ep.Validation.HasErrors() {
		ep.Validation.Keep()
		ep.FlashParams()
		return ep.Redirect(ExamPaper.Publish)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Publish)
	}
	defer manager.Close()

	examinees, err := manager.GetAllExaminee()
	if err != nil {
		log.Println(err)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Publish)
	}

	examPaper, err := manager.GetExamPaperByTitle(exmpaperTitle)
	if err != nil {
		log.Println(err)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Redirect(ExamPaper.Publish)
	}

	for _, examinee := range examinees {
		examPaper.Status = models.UnFinished
		models.ChaosExamPaper(&examPaper)
		examinee.ExamPaper = examPaper

		err := manager.UpdateExaminee(&examinee)
		if err != nil {
			log.Println(err)
			ep.Response.Status = 500
			ep.Flash.Error(err.Error())
			return ep.Redirect(ExamPaper.Publish)
		}
	}

	ep.Flash.Success("考试成功发布")

	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Redirect(ExamPaper.Publish)
}

func (ep ExamPaper) Score(idCard string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}
	defer manager.Close()

	examinee, e := manager.GetExamineeByIDCard(idCard)

	if e != nil {
		log.Println(e)
		ep.Response.Status = 500
		ep.Flash.Error(err.Error())
		return ep.Render()
	}

	log.Println("查到考生信息: ", examinee)

	ep.ViewArgs["examinee"] = examinee
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) QueryScore() revel.Result {
	ep.ViewArgs["adminIDCard"] = ep.Session["adminIDCard"]
	ep.ViewArgs["adminName"] = ep.Session["adminName"]

	return ep.Render()
}

func (ep ExamPaper) PostQueryScore(examineeIDCard string) revel.Result {
	examineeIDCard = strings.TrimSpace(examineeIDCard)

	ep.Validation.Required(examineeIDCard).Message("请输入身份证号码")
	ep.Validation.Length(examineeIDCard, 18).Message("身份证号有误")

	if ep.Validation.HasErrors() {
		ep.Validation.Keep()
		ep.FlashParams()
		return ep.Redirect(ExamPaper.QueryScore)
	}

	return ep.Redirect("/ExamPaper/Score/%s", examineeIDCard)
	//return ep.Redirect(ExamPaper.Score("huguangyue"))
}
