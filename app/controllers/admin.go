package controllers

import (
	"ExamSystem/app/models"
	"log"

	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
}

func (this Admin) Index() revel.Result {
	this.RenderArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.RenderArgs["adminName"] = this.Session["adminName"]
	
	return this.Render()
}

func (this Admin) SignUp() revel.Result {
	if this.Session["SignUpStatus"] == "true" {
		this.RenderArgs["SignUpStatus"] = true
		this.Session["SignUpStatus"] = "false"
	}
	
	return this.Render()
}

func (this Admin) PostSignUp(signUpAdmin *models.SignUpAdmin) revel.Result {
	this.Validation.Required(signUpAdmin.Name).Message("请输入管理员姓名")
	this.Validation.Required(signUpAdmin.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signUpAdmin.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signUpAdmin.Password).Message("请输入密码")
	this.Validation.Required(signUpAdmin.ConfirmPassword).Message("确认密码不能为空")
	this.Validation.MinSize(signUpAdmin.Password, 6).Message("密码长度不短于6位")
	this.Validation.Required(signUpAdmin.ConfirmPassword == signUpAdmin.Password).Message("两次输入的密码不一致")

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

	err = manager.AdminSignUp(signUpAdmin)
	if err != nil {
		this.Validation.Clear()

		// 添加错误信息，显示在页面的身份证下面
		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "signUpAdmin.IDCard"
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Admin.SignUp)
	}

	this.Session["SignUpStatus"] = "true"
	log.Println("注册成功！")
	log.Println(signUpAdmin)

	return this.Redirect(Admin.SignUp)
}

func (this Admin) SignIn() revel.Result {
	return this.Render()
}

func (this Admin) PostSignIn(signInAdmin *models.SignInAdmin) revel.Result {
	this.Validation.Required(signInAdmin.IDCard).Message("请输入身份证号码")
	this.Validation.Length(signInAdmin.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(signInAdmin.Password).Message("请输入密码")

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		log.Println(this.Validation.Errors)		
		return this.Redirect(Admin.SignIn)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	var a *models.Admin
	a, err = manager.AdminSignIn(signInAdmin)

	if err != nil {
		this.Validation.Clear()

		// 添加错误提示信息，显示在页面的该管理员身份证/密码下面
		var e revel.ValidationError
		if err.Error() == "该管理员不存在" {
			e.Key = "signInAdmin.IDCard"
		} else {
			e.Key = "signInAdmin.Password"
		}
		e.Message = err.Error()
		this.Validation.Errors = append(this.Validation.Errors, &e)

		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(Admin.SignIn)
	}

	this.Session["adminIDCard"] = a.IDCard
	this.Session["adminName"] = a.Name

	this.RenderArgs["adminName"] = a.Name
	log.Println("登录成功: ", a)

	return this.Redirect(Admin.Index)
}