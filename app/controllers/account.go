package controllers

import (
	"ExamSystem/app/models"
	"log"

	"github.com/revel/revel"
)

type Account struct {
	*revel.Controller
}

func (this Account) Index() revel.Result {
	return this.Render()
}

func (this Account) SignUp() revel.Result {
	return this.Render()
}

func (this Account) PostSignUp(mu *models.MockUser) revel.Result {
	this.Validation.Required(mu.Name).Message("请输入考生姓名")
	this.Validation.Required(mu.IDCard).Message("请输入身份证号码")
	this.Validation.MinSize(mu.IDCard, 18).Message("身份证号有误")
	this.Validation.Required(mu.Password).Message("请输入密码")
	this.Validation.Required(mu.ConfirmPassword).Message("确认密码不能为空")
	this.Validation.MinSize(mu.Password, 6).Message("密码长度不短于6位")
	this.Validation.Required(mu.ConfirmPassword == mu.Password).Message("两次输入的密码不一致")

	log.Println(mu)
	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Account).SignUp)
	}
	return this.Redirect((*Account).SignUp)

	//	manager, err := models.NewDbManager()
	//	if err != nil {
	//		this.Response.Status = 500
	//		return this.RenderError(err)
	//	}
	//	defer manager.Close()

	//	err = manager.RegisterUser(mu)
	//	if err != nil {
	//		this.Validation.Clear()

	//		// 添加错误信息，显示在页面的用户名下面
	//		var e revel.ValidationError
	//		e.Message = err.Error()
	//		e.Key = "mu.NickName"
	//		this.Validation.Errors = append(this.Validation.Errors, &e)

	//		this.Validation.Keep()
	//		this.FlashParams()
	//		return this.Redirect((*Account).Register)
	//	}

	//	return this.Redirect((*Account).RegisterSuccessful)
}

func (this Account) SignIn() revel.Result {
	return this.Render()
}

func (this Account) SignOut() revel.Result {
	return this.Render()
}
