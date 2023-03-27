package controllers

import (
	"examsystem/app/models"
	"log"
	"strings"

	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
}

func (this Admin) Index() revel.Result {
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this Admin) SignUp() revel.Result {
	this.ViewArgs["adminIDCard"] = this.Session["adminIDCard"]
	this.ViewArgs["adminName"] = this.Session["adminName"]

	return this.Render()
}

func (this Admin) PostSignUp(signUpAdmin *models.SignUpAdmin) revel.Result {
	signUpAdmin.Name = strings.TrimSpace(signUpAdmin.Name)
	signUpAdmin.IDCard = strings.TrimSpace(signUpAdmin.IDCard)
	signUpAdmin.Password = strings.TrimSpace(signUpAdmin.Password)
	signUpAdmin.ConfirmPassword = strings.TrimSpace(signUpAdmin.ConfirmPassword)

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
		this.Flash.Error(err.Error())
		return this.Redirect((*Examinee).SignUp)
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
		this.Flash.Error(err.Error())
		return this.Redirect(Admin.SignUp)
	}

	log.Println("注册成功：" + signUpAdmin.Name)
	this.Flash.Success("注册成功：" + signUpAdmin.Name)

	return this.Redirect(Admin.SignUp)
}

func (this Admin) SignIn() revel.Result {
	return this.Render()
}

func (this Admin) PostSignIn(signInAdmin *models.SignInAdmin) revel.Result {
	signInAdmin.IDCard = strings.TrimSpace(signInAdmin.IDCard)
	signInAdmin.Password = strings.TrimSpace(signInAdmin.Password)

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
		this.Flash.Error(err.Error())
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
		this.Flash.Error(err.Error())
		return this.Redirect(Admin.SignIn)
	}

	this.Session["adminIDCard"] = a.IDCard
	this.Session["adminName"] = a.Name
	this.Session["administrator"] = "true"

	this.ViewArgs["adminIDCard"] = a.IDCard
	this.ViewArgs["adminName"] = a.Name
	log.Println("登录成功: ", a)

	return this.Redirect(Admin.Index)
}

func (this Admin) SignOut() revel.Result {
	for k := range this.Session {
		delete(this.Session, k)
	}
	return this.Redirect(Admin.SignIn)
}
