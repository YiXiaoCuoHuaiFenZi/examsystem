package controllers

import (
	"encoding/csv"
	"examsystem/app/models"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

type Examinee struct {
	*revel.Controller
}

func (ex Examinee) Index() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.RenderError(err)
	}
	defer manager.Close()

	idCard := ex.Session["examineeIDCard"].(string)
	examinee, e := manager.GetExamineeByIDCard(idCard)

	if e != nil {
		ex.Response.Status = 500
		ex.Flash.Error(e.Error())
		return ex.RenderError(e)
	}

	ex.ViewArgs["examinee"] = examinee
	ex.ViewArgs["examineeIDCard"] = idCard
	ex.ViewArgs["examineeName"] = ex.Session["examineeName"]

	return ex.Render()
}

func (ex Examinee) Info() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		return ex.Render(err)
	}
	defer manager.Close()

	examinees, e := manager.GetAllExaminee()
	if e != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Render(e)
	}

	ex.ViewArgs["examinees"] = examinees
	ex.ViewArgs["adminIDCard"] = ex.Session["adminIDCard"]
	ex.ViewArgs["adminName"] = ex.Session["adminName"]

	return ex.Render()
}

func (ex Examinee) SignUp() revel.Result {
	ex.ViewArgs["batch"] = ex.Session["batch"]
	ex.ViewArgs["adminIDCard"] = ex.Session["adminIDCard"]
	ex.ViewArgs["adminName"] = ex.Session["adminName"]

	return ex.Render()
}

func (ex Examinee) PostSignUp(signUpExaminee *models.SignUpExaminee) revel.Result {
	signUpExaminee.Name = strings.TrimSpace(signUpExaminee.Name)
	signUpExaminee.IDCard = strings.TrimSpace(signUpExaminee.IDCard)
	signUpExaminee.Password = strings.TrimSpace(signUpExaminee.Password)
	signUpExaminee.ConfirmPassword = strings.TrimSpace(signUpExaminee.ConfirmPassword)

	ex.Validation.Required(signUpExaminee.Name).Message("请输入考生姓名")
	ex.Validation.Required(signUpExaminee.IDCard).Message("请输入身份证号码")
	ex.Validation.Length(signUpExaminee.IDCard, 18).Message("身份证号有误")
	ex.Validation.Required(signUpExaminee.Password).Message("请输入密码")
	ex.Validation.Required(signUpExaminee.ConfirmPassword).Message("确认密码不能为空")
	ex.Validation.MinSize(signUpExaminee.Password, 6).Message("密码长度不短于6位")
	ex.Validation.Required(signUpExaminee.ConfirmPassword == signUpExaminee.Password).Message("两次输入的密码不一致")

	if ex.Validation.HasErrors() {
		ex.Validation.Keep()
		ex.FlashParams()
		return ex.Redirect(Examinee.SignUp)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.SignUp)
	}
	defer manager.Close()

	err = manager.SignUp(signUpExaminee)
	if err != nil {
		ex.Validation.Clear()

		// 添加错误信息，显示在页面的身份证下面
		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "signUpExaminee.IDCard"
		ex.Validation.Errors = append(ex.Validation.Errors, &e)

		ex.Validation.Keep()
		ex.FlashParams()
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.SignUp)
	}

	log.Println("注册成功：" + signUpExaminee.Name)
	ex.Flash.Success("注册成功：" + signUpExaminee.Name)

	return ex.Redirect(Examinee.SignUp)
}

func (ex Examinee) BatchSignUp() revel.Result {
	return ex.Redirect(Examinee.SignUp)
}

func (ex Examinee) PostBatchSignUp(CSVFile *os.File) revel.Result {
	// TODO csv文件默认是ascII编码， 需要进行处理
	// 暂时强制要求手动转换为utf8
	reader := csv.NewReader(CSVFile)
	defer CSVFile.Close()

	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.SignUp)
	}
	defer manager.Close()

	var i = 0
	var errorMsg = ""
	var successMsg = ""
	for {
		lineArr, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			ex.Flash.Error(err.Error())
			return ex.Redirect(Examinee.SignUp)
		}

		//忽略第一行表头
		i += 1
		if i == 1 {
			continue
		}

		var e models.SignUpExaminee
		e.Name = lineArr[0]
		e.IDCard = lineArr[1]
		e.Gender = lineArr[2]
		e.Mobile = lineArr[3]
		// 密码为身份证后六位
		e.Password = e.IDCard[len(e.IDCard)-6:]
		e.ConfirmPassword = e.Password

		err = manager.SignUp(&e)
		if err != nil {
			m := err.Error() + "：" + e.IDCard + "\n"
			errorMsg += m
			log.Println(m)
		} else {
			successMsg += "注册成功：" + e.IDCard + "\n"
			log.Println("注册成功：", e.IDCard)
		}
	}

	ex.Flash.Success(successMsg + errorMsg)
	ex.Session["batch"] = "true"
	return ex.Redirect(Examinee.SignUp)
}

func (ex Examinee) SignIn() revel.Result {
	ex.ViewArgs["adminIDCard"] = ex.Session["adminIDCard"]
	ex.ViewArgs["adminName"] = ex.Session["adminName"]
	ex.ViewArgs["examineeIDCard"] = ex.Session["examineeIDCard"]
	ex.ViewArgs["examineeName"] = ex.Session["examineeName"]

	return ex.Render()
}

func (ex Examinee) PostSignIn(signInExaminee *models.SignInExaminee) revel.Result {
	signInExaminee.IDCard = strings.TrimSpace(signInExaminee.IDCard)
	signInExaminee.Password = strings.TrimSpace(signInExaminee.Password)

	ex.Validation.Required(signInExaminee.IDCard).Message("请输入身份证号码")
	ex.Validation.Length(signInExaminee.IDCard, 18).Message("身份证号有误")
	ex.Validation.Required(signInExaminee.Password).Message("请输入密码")

	if ex.Validation.HasErrors() {
		ex.Validation.Keep()
		ex.FlashParams()
		return ex.Redirect(Examinee.SignIn)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.SignIn)
	}
	defer manager.Close()

	var e *models.Examinee
	e, err = manager.SignIn(signInExaminee)

	if err != nil {
		ex.Validation.Clear()

		// 添加错误提示信息，显示在页面的用户名/密码下面
		var e revel.ValidationError
		if err.Error() == "该考生不存在" {
			e.Key = "signInExaminee.IDCard"
		} else {
			e.Key = "signInExaminee.Password"
		}
		e.Message = err.Error()
		ex.Validation.Errors = append(ex.Validation.Errors, &e)

		ex.Validation.Keep()
		ex.FlashParams()
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.SignIn)
	}

	ex.Session["examineeIDCard"] = e.IDCard
	ex.Session["examineeName"] = e.Name
	ex.Session["examinee"] = "true"

	ex.ViewArgs["examineeIDCard"] = e.IDCard
	ex.ViewArgs["examineeName"] = e.Name
	log.Println("登录成功: ", e.Name)

	return ex.Redirect(Examinee.Index)
}

func (ex Examinee) Exam(examPaperTitle string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Render()
	}
	defer manager.Close()

	examinee, err := manager.GetExamineeByIDCard(ex.Session["examineeIDCard"].(string))
	if err != nil {
		log.Println(err)
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Render()
	}

	if examinee.ExamPaper.Status == models.Done {
		return ex.Redirect("/Examinee/ExamResult?idCard=%s&title=%s", ex.Session["examineeIDCard"], examPaperTitle)
	}

	scCount := len(examinee.ExamPaper.SC)
	mcCount := len(examinee.ExamPaper.MC)
	tfCount := len(examinee.ExamPaper.TF)

	scws := make([]models.SCWithPage, 0)
	mcws := make([]models.MCWithPage, 0)
	tfws := make([]models.TFWithPage, 0)

	npp, left := 10, 0 // npp: number per page
	for index, item := range examinee.ExamPaper.SC {
		scws = append(scws, models.SCWithPage{Page: index/npp + 1, SC: item})
	}
	left = scCount % npp

	for index, item := range examinee.ExamPaper.MC {
		mcws = append(mcws, models.MCWithPage{Page: scCount/npp + (left+index)/npp + 1, MC: item})
	}
	left = (scCount + mcCount) % npp

	for index, item := range examinee.ExamPaper.TF {
		tfws = append(tfws, models.TFWithPage{Page: (scCount+mcCount)/npp + (left+index)/npp + 1, TF: item})
	}

	p := 0
	if (scCount+mcCount+tfCount)%npp > 0 {
		p = (scCount+mcCount+tfCount)/npp + 1
	} else {
		p = (scCount + mcCount + tfCount) / npp
	}

	pages := make([]int, 0)
	for i := 1; i <= p; i++ {
		pages = append(pages, i)
	}

	ex.ViewArgs["exam"] = "true"
	// TODO 支持多试卷，根据试卷标题查询得到要考试的试卷
	ex.ViewArgs["scws"] = scws
	ex.ViewArgs["mcws"] = mcws
	ex.ViewArgs["tfws"] = tfws
	ex.ViewArgs["pages"] = pages

	ex.ViewArgs["scCount"] = scCount
	ex.ViewArgs["mcCount"] = mcCount
	ex.ViewArgs["tfCount"] = tfCount
	ex.ViewArgs["examPaper"] = examinee.ExamPaper
	ex.ViewArgs["examineeIDCard"] = ex.Session["examineeIDCard"]
	ex.ViewArgs["examineeName"] = ex.Session["examineeName"]

	return ex.Render()
}

func (ex Examinee) PostExam() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.Exam)
	}
	defer manager.Close()

	examineeIDCard := ex.Session["examineeIDCard"].(string)
	examinee, err := manager.GetExamineeByIDCard(examineeIDCard)
	if err != nil {
		log.Println(err)
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.Exam)
	}

	leftTime := ex.Params.Form.Get("leftTime")
	if leftTime != "" {
		lt, err := strconv.ParseInt(leftTime, 10, 64)
		if err != nil {
			log.Println("转换剩余时间出错")
			return nil
		}
		examinee.ExamPaper.LeftTime = lt
		examinee.ExamPaper.Status = models.Doing
		log.Println("自动保存 " + examinee.Name + " 的试卷，考试剩余时间：" + leftTime + "毫秒")
	} else {
		examinee.ExamPaper.Status = models.Done
	}

	// 页面上显示的name值已经增加了1，所以这里需要加1，以将其对应起来
	for index, _ := range examinee.ExamPaper.SC {
		answer := ex.Params.Form.Get("sc_" + strconv.Itoa(index+1) + "_answer")
		examinee.ExamPaper.SC[index].ActualAnswer = answer
	}

	for index, _ := range examinee.ExamPaper.MC {
		answers := ex.Params.Form["mc_"+strconv.Itoa(index+1)+"_answers[]"]
		examinee.ExamPaper.MC[index].ActualAnswer = answers
	}

	for index, _ := range examinee.ExamPaper.TF {
		answer := ex.Params.Form.Get("tf_" + strconv.Itoa(index+1) + "_answer")
		examinee.ExamPaper.TF[index].ActualAnswer = answer
	}
	models.MarkExamPaper(&examinee.ExamPaper)

	err = manager.UpdateExamPaper(examineeIDCard, examinee.ExamPaper)
	if err != nil {
		log.Println(err)
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Redirect(Examinee.Exam)
	}

	ex.Flash.Success("成功交卷")
	ex.ViewArgs["examineeIDCard"] = ex.Session["examineeIDCard"]
	ex.ViewArgs["examineeName"] = ex.Session["examineeName"]
	return ex.Redirect(Examinee.Index)
}

func (ex Examinee) ExamResult(idCard, title string) revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Render()
	}
	defer manager.Close()

	examinee, err := manager.GetExamineeByIDCard(idCard)
	if err != nil {
		log.Println(err)
		ex.Response.Status = 500
		ex.Flash.Error(err.Error())
		return ex.Render()
	}

	ex.ViewArgs["scCount"] = len(examinee.ExamPaper.SC)
	ex.ViewArgs["mcCount"] = len(examinee.ExamPaper.MC)
	ex.ViewArgs["tfCount"] = len(examinee.ExamPaper.TF)
	ex.ViewArgs["examPaper"] = examinee.ExamPaper

	ex.ViewArgs["adminIDCard"] = ex.Session["adminIDCard"]
	ex.ViewArgs["adminName"] = ex.Session["adminName"]
	ex.ViewArgs["examineeIDCard"] = ex.Session["examineeIDCard"]
	ex.ViewArgs["examineeName"] = ex.Session["examineeName"]

	return ex.Render()
}

func (ex Examinee) SignOut() revel.Result {
	for k := range ex.Session {
		delete(ex.Session, k)
	}
	return ex.Redirect(Examinee.SignIn)
}
