package controllers

import (
	"ExamSystem/app/models"
	"log"

	"github.com/revel/revel"
)

type Examinee struct {
	*revel.Controller
}

func (this Examinee) Index() revel.Result {
	return this.Render()
}

func (this Examinee) Info() revel.Result {
	return this.Render()
}

func (this Examinee) SignUp() revel.Result {
	if this.Session["SignUpStatus"] == "true" {
		this.RenderArgs["SignUpStatus"] = true
		this.Session["SignUpStatus"] = "false"
	}
	return this.Render()
}

func (this Examinee) PostSignUp(signUpUser *models.SignUpUser) revel.Result {
	this.Validation.Required(signUpUser.Name).Message("请输入考生姓名")
	this.Validation.Required(signUpUser.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signUpUser.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signUpUser.Password).Message("请输入密码")
	this.Validation.Required(signUpUser.ConfirmPassword).Message("确认密码不能为空")
	this.Validation.MinSize(signUpUser.Password, 6).Message("密码长度不短于6位")
	this.Validation.Required(signUpUser.ConfirmPassword == signUpUser.Password).Message("两次输入的密码不一致")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Examinee).SignUp)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.SignUp(signUpUser)
	if err != nil {
		this.Validation.Clear()

		// 添加错误信息，显示在页面的身份证下面
		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "signUpUser.IDCard"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Examinee).SignUp)
	}

	this.Session["SignUpStatus"] = "true"
	log.Println("注册成功！")
	log.Println(signUpUser)

	return this.Redirect((*Examinee).SignUp)
}

func (this Examinee) SignIn() revel.Result {
	return this.Render()
}

func (this Examinee) PostSignIn(signInUser *models.SignInUser) revel.Result {
	this.Validation.Required(signInUser.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signInUser.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signInUser.Password).Message("请输入密码")

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

	var u *models.User
	u, err = manager.SignIn(signInUser)

	if err != nil {
		this.Validation.Clear()

		// 添加错误提示信息，显示在页面的用户名/密码下面
		var e revel.ValidationError
		if err.Error() == "该用户不存在" {
			e.Key = "signInUser.IDCard"
		} else {
			e.Key = "signInUser.Password"
		}
		e.Message = err.Error()
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Examinee).SignIn)
	}

	this.Session["userIDCard"] = u.IDCard
	this.Session["userName"] = u.Name

	this.RenderArgs["userName"] = u.Name
	log.Println("登录成功: ", u)

	return this.Redirect((*Examinee).Index)
}

func (this Examinee) Exam() revel.Result {
	return this.Render()
}

func (this Examinee) SignOut() revel.Result {
	return this.Render()
}