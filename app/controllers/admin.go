package controllers

import (
	"ExamSystem/app/models"

	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
}

func (this Admin) Index() revel.Result {
	return this.Render()
}

func (this Admin) SignIn() revel.Result {
	return this.Render()
}

func (this Admin) PostSignIn(signInAdmin *models.SignInAdmin) revel.Result {
	//	this.Validation.Required(signInUser.IDCard).Message("请输入身份证号码")
	//	this.Validation.Length(signInUser.IDCard, 18).Message("身份证号有误")
	//	this.Validation.Required(signInUser.Password).Message("请输入密码")

	//	if this.Validation.HasErrors() {
	//		this.Validation.Keep()
	//		this.FlashParams()
	//		return this.Redirect((*Account).SignIn)
	//	}

	//	manager, err := models.NewDBManager()
	//	if err != nil {
	//		this.Response.Status = 500
	//		return this.RenderError(err)
	//	}
	//	defer manager.Close()

	//	var u *models.User
	//	u, err = manager.SignIn(signInUser)

	//	if err != nil {
	//		this.Validation.Clear()

	//		// 添加错误提示信息，显示在页面的用户名/密码下面
	//		var e revel.ValidationError
	//		if err.Error() == "该用户不存在" {
	//			e.Key = "signInUser.IDCard"
	//		} else {
	//			e.Key = "signInUser.Password"
	//		}
	//		e.Message = err.Error()
	//		this.Validation.Errors = append(this.Validation.Errors, &e)

	//		this.Validation.Keep()
	//		this.FlashParams()
	//		return this.Redirect((*Account).SignIn)
	//	}

	//	this.Session["userIDCard"] = u.IDCard
	//	this.Session["userName"] = u.Name

	//	this.RenderArgs["userName"] = u.Name
	//	log.Println("登录成功: ", u)

	return this.Redirect(Admin.Index)
}
