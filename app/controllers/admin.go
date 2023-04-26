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

func (ad Admin) Index() revel.Result {
	ad.ViewArgs["adminIDCard"] = ad.Session["adminIDCard"]
	ad.ViewArgs["adminName"] = ad.Session["adminName"]

	return ad.Render()
}

func (ad Admin) SignUp() revel.Result {
	ad.ViewArgs["adminIDCard"] = ad.Session["adminIDCard"]
	ad.ViewArgs["adminName"] = ad.Session["adminName"]

	return ad.Render()
}

func (ad Admin) PostSignUp(signUpAdmin *models.SignUpAdmin) revel.Result {
	signUpAdmin.Name = strings.TrimSpace(signUpAdmin.Name)
	signUpAdmin.IDCard = strings.TrimSpace(signUpAdmin.IDCard)
	signUpAdmin.Password = strings.TrimSpace(signUpAdmin.Password)
	signUpAdmin.ConfirmPassword = strings.TrimSpace(signUpAdmin.ConfirmPassword)

	ad.Validation.Required(signUpAdmin.Name).Message("请输入管理员姓名")
	ad.Validation.Required(signUpAdmin.IDCard).Message("请输入身份证号码")
	ad.Validation.Length(signUpAdmin.IDCard, 18).Message("身份证号有误")
	ad.Validation.Required(signUpAdmin.Password).Message("请输入密码")
	ad.Validation.Required(signUpAdmin.ConfirmPassword).Message("确认密码不能为空")
	ad.Validation.MinSize(signUpAdmin.Password, 6).Message("密码长度不短于6位")
	ad.Validation.Required(signUpAdmin.ConfirmPassword == signUpAdmin.Password).Message("两次输入的密码不一致")

	if ad.Validation.HasErrors() {
		ad.Validation.Keep()
		ad.FlashParams()
		return ad.Redirect((*Examinee).SignUp)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ad.Response.Status = 500
		ad.Flash.Error(err.Error())
		return ad.Redirect((*Examinee).SignUp)
	}
	defer manager.Close()

	err = manager.AdminSignUp(signUpAdmin)
	if err != nil {
		ad.Validation.Clear()

		// 添加错误信息，显示在页面的身份证下面
		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "signUpAdmin.IDCard"
		ad.Validation.Errors = append(ad.Validation.Errors, &e)

		ad.Validation.Keep()
		ad.FlashParams()
		ad.Flash.Error(err.Error())
		return ad.Redirect(Admin.SignUp)
	}

	log.Println("注册成功：" + signUpAdmin.Name)
	ad.Flash.Success("注册成功：" + signUpAdmin.Name)

	return ad.Redirect(Admin.SignUp)
}

func (ad Admin) SignIn() revel.Result {
	return ad.Render()
}

func (ad Admin) PostSignIn(signInAdmin *models.SignInAdmin) revel.Result {
	signInAdmin.IDCard = strings.TrimSpace(signInAdmin.IDCard)
	signInAdmin.Password = strings.TrimSpace(signInAdmin.Password)

	ad.Validation.Required(signInAdmin.IDCard).Message("请输入身份证号码")
	ad.Validation.Length(signInAdmin.IDCard, 18).Message("身份证号有误")
	ad.Validation.Required(signInAdmin.Password).Message("请输入密码")

	if ad.Validation.HasErrors() {
		ad.Validation.Keep()
		ad.FlashParams()
		log.Println(ad.Validation.Errors)
		return ad.Redirect(Admin.SignIn)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		ad.Response.Status = 500
		ad.Flash.Error(err.Error())
		return ad.RenderError(err)
	}
	defer manager.Close()

	var a *models.Admin
	a, err = manager.AdminSignIn(signInAdmin)

	if err != nil {
		ad.Validation.Clear()

		// 添加错误提示信息，显示在页面的该管理员身份证/密码下面
		var e revel.ValidationError
		if err.Error() == "该管理员不存在" {
			e.Key = "signInAdmin.IDCard"
		} else {
			e.Key = "signInAdmin.Password"
		}
		e.Message = err.Error()
		ad.Validation.Errors = append(ad.Validation.Errors, &e)

		ad.Validation.Keep()
		ad.FlashParams()
		ad.Flash.Error(err.Error())
		return ad.Redirect(Admin.SignIn)
	}

	ad.Session["adminIDCard"] = a.IDCard
	ad.Session["adminName"] = a.Name
	ad.Session["administrator"] = "true"

	ad.ViewArgs["adminIDCard"] = a.IDCard
	ad.ViewArgs["adminName"] = a.Name
	log.Println("登录成功: ", a)

	return ad.Redirect(Admin.Index)
}

func (ad Admin) SignOut() revel.Result {
	for k := range ad.Session {
		delete(ad.Session, k)
	}
	return ad.Redirect(Admin.SignIn)
}
