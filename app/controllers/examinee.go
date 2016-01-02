package controllers

import (
	"ExamSystem/app/models"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/revel/revel"
)

type Examinee struct {
	*revel.Controller
}

func (this Examinee) Index() revel.Result {
	this.RenderArgs["examnineeIDCard"] = this.Session["examnineeIDCard"]
	this.RenderArgs["examnineeName"] = this.Session["examnineeName"]

	return this.Render()
}

func (this Examinee) Info() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this Examinee) SignUp() revel.Result {
	if this.Session["SignUpStatus"] == "true" {
		this.RenderArgs["SignUpStatus"] = true
		this.Session["SignUpStatus"] = "false"
	}

	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this Examinee) PostSignUp(signUpExaminee *models.SignUpExaminee) revel.Result {
	this.Validation.Required(signUpExaminee.Name).Message("请输入考生姓名")
	this.Validation.Required(signUpExaminee.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signUpExaminee.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signUpExaminee.Password).Message("请输入密码")
	this.Validation.Required(signUpExaminee.ConfirmPassword).Message("确认密码不能为空")
	this.Validation.MinSize(signUpExaminee.Password, 6).Message("密码长度不短于6位")
	this.Validation.Required(signUpExaminee.ConfirmPassword == signUpExaminee.Password).Message("两次输入的密码不一致")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Examinee.SignUp)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.SignUp(signUpExaminee)
	if err != nil {
		this.Validation.Clear()

		// 添加错误信息，显示在页面的身份证下面
		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "signUpExaminee.IDCard"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Examinee.SignUp)
	}

	this.Session["SignUpStatus"] = "true"
	log.Println("注册成功！")
	log.Println(signUpExaminee)

	return this.Redirect((*Examinee).SignUp)
}

func (this Examinee) BatchSignUp() revel.Result {
	return this.Redirect(Examinee.SignUp)
}

func (this Examinee) PostBatchSignUp(CSVFile *os.File) revel.Result {
	// TODO csv文件默认是ascII编码， 需要进行处理
	// 暂时强制要求手动转换为utf8
	reader := csv.NewReader(CSVFile)
	defer CSVFile.Close()

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
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
			return this.RenderError(err)
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
		// 密码为身份证后六位
		e.Password = e.IDCard[len(e.IDCard)-6:]
		e.ConfirmPassword = e.Password

		err = manager.SignUp(&e)
		if err != nil {
			m := err.Error() + "：" + e.IDCard + "  <br>"
			errorMsg += m
			log.Println(m)
		} else {
			successMsg += "注册成功：" + e.IDCard + "  <br>"
			log.Println("注册成功：", e)
		}
	}

	this.Flash.Error(successMsg + errorMsg)
	//	if errorMsg != "" {
	//		this.Flash.Error(successMsg+errorMsg)
	//	}
	//	if successMsg != "" {
	//		this.Flash.Success("注册成功：", successMsg)
	//	}
	//this.Session["SignUpStatus"] = "true"
	return this.Redirect(Examinee.SignUp)
}

func (this Examinee) SignIn() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	this.RenderArgs["examnineeIDCard"] = this.Session["examnineeIDCard"]
	this.RenderArgs["examnineeName"] = this.Session["examnineeName"]

	return this.Render()
}

func (this Examinee) PostSignIn(signInExaminee *models.SignInExaminee) revel.Result {
	this.Validation.Required(signInExaminee.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signInExaminee.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signInExaminee.Password).Message("请输入密码")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Examinee).SignIn)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	var e *models.Examinee
	e, err = manager.SignIn(signInExaminee)

	if err != nil {
		this.Validation.Clear()

		// 添加错误提示信息，显示在页面的用户名/密码下面
		var e revel.ValidationError
		if err.Error() == "该考生不存在" {
			e.Key = "signInExaminee.IDCard"
		} else {
			e.Key = "signInExaminee.Password"
		}
		e.Message = err.Error()
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Examinee).SignIn)
	}

	this.Session["examnineeIDCard"] = e.IDCard
	this.Session["examnineeName"] = e.Name
	this.Session["examinee"] = "true"

	this.RenderArgs["examnineeIDCard"] = e.IDCard
	this.RenderArgs["examnineeName"] = e.Name
	log.Println("登录成功: ", e)

	return this.Redirect(Examinee.Index)
}

func (this Examinee) Exam() revel.Result {
	this.RenderArgs["examnineeIDCard"] = this.Session["examnineeIDCard"]
	this.RenderArgs["examnineeName"] = this.Session["examnineeName"]

	return this.Render()
}

func (this Examinee) SignOut() revel.Result {
	for k := range this.Session {
		delete(this.Session, k)
	}
	return this.Redirect(Examinee.SignIn)
}
