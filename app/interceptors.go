package app

import (
	"examsystem/app/controllers"
	"examsystem/app/models"
	"log"

	"github.com/revel/revel"
)

func InitDB() {
	log.Println("初始化数据库")
	revel.Config.SetSection("administrator")
	idcard, found := revel.Config.String("id_card")
	if !found {
		log.Fatal("配置文件app.conf没有找到超级管理员administrator的身份证号码id_card")
	}

	password, found := revel.Config.String("password")
	if !found {
		log.Fatal("配置文件app.conf没有找到超级管理员administrator的密码配置password")
	}

	manager, err := models.NewDBManager()
	if err != nil {
		log.Fatal(err)
	}

	err = manager.CreateSuperAdministrator(idcard, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("初始化数据库完成")
}

func beforeAdminController(this *revel.Controller) revel.Result {
	// 设置游客的访问权限
	if this.Session["administrator"] == nil && this.Session["examinee"] == nil {
		if this.Action == "Admin.SignIn" || this.Action == "Admin.PostSignIn" {
			return nil
		} else {
			this.Flash.Error("您需要登录才能访问该页面。")
			log.Println("游客访问了 " + this.Action + " 页面，已自动跳转到登陆页面")
			return this.Redirect(controllers.Admin.SignIn)
		}
	}

	// 设置考生的访问权限
	if this.Session["examinee"] == "true" {
		this.Flash.Error("您没有访问该页面的权限，已自动跳转到首页。")
		log.Println("考生访问了禁止的页面: " + this.Action + "，已自动跳转到考生首页")
		return this.Redirect(controllers.Examinee.Index)
	}

	// 设置管理员的访问权限
	if this.Session["administrator"] == "true" {
		if this.Action == "Admin.SignIn" || this.Action == "Admin.PostSignIn" {
			this.Flash.Error("您已经登录，请先注销再登录。")
			log.Println("管理员访问了登录页面，已自动跳转到管理员首页")
			return this.Redirect(controllers.Admin.Index)
		} else {
			log.Println("管理员 " + this.Session["adminName"].(string) + " 访问了 " + this.Action + " 页面")
			return nil
		}
	}

	return this.Redirect(controllers.Admin.SignIn)
}

func beforeExamineeController(this *revel.Controller) revel.Result {
	// 设置游客的访问权限
	if this.Session["administrator"] == nil && this.Session["examinee"] == nil {
		if this.Action == "Examinee.SignIn" || this.Action == "Examinee.PostSignIn" {
			return nil
		} else {
			this.Flash.Error("您需要登录才能访问该页面。")
			log.Println("游客访问了 " + this.Action + " 页面，已自动跳转到登陆页面")
			return this.Redirect(controllers.Examinee.SignIn)
		}
	}

	// 设置考生的访问权限
	if this.Session["examinee"] == "true" {
		if this.Action == "Examinee.Exam" || this.Action == "Examinee.PostExam" ||
			this.Action == "Examinee.Index" || this.Action == "Examinee.SignOut" ||
			this.Action == "Examinee.ExamResult" {
			return nil
		} else if this.Action == "Examinee.SignIn" || this.Action == "Examinee.PostSignIn" {
			this.Flash.Error("您已经登录，请先注销再登录。")
			log.Println("考生访问了登录页面，已自动跳转到考生首页")
			return this.Redirect(controllers.Examinee.Index)
		} else {
			this.Flash.Error("您没有访问该页面的权限，已自动跳转到首页。")
			log.Println("考生访问了禁止的页面: " + this.Action + "，已自动跳转到考生首页")
			return this.Redirect(controllers.Examinee.Index)
		}
	}

	// 设置管理员的访问权限
	if this.Session["administrator"] == "true" {
		if this.Action == "Examinee.SignIn" || this.Action == "Examinee.PostSignIn" {
			this.Flash.Error("您已经登录，请先注销再登录。")
			log.Println("管理员访问了考生登录页面，已自动跳转到管理员首页")
			return this.Redirect(controllers.Admin.Index)
		} else {
			log.Println("管理员 " + this.Session["adminName"].(string) + " 访问了 " + this.Action + " 页面")
			return nil
		}
	}

	return this.Redirect(controllers.Examinee.SignIn)
}

func beforeAboutController(this *revel.Controller) revel.Result {
	if this.Session["administrator"] == "true" {
		return nil
	} else {
		this.Flash.Error("您需要登录才能访问该页面。")
		log.Println("非管理员访问了 " + this.Action + " 页面，已自动跳转到登陆页面")
		return this.Redirect(controllers.Admin.SignIn)
	}
}

func beforeExamPaperController(this *revel.Controller) revel.Result {
	if this.Session["administrator"] == "true" {
		return nil
	} else {
		this.Flash.Error("您需要登录才能访问该页面。")
		log.Println("非管理员访问了 " + this.Action + " 页面，已自动跳转到登陆页面")
		return this.Redirect(controllers.Admin.SignIn)
	}
}

func beforeQuestionController(this *revel.Controller) revel.Result {
	if this.Session["administrator"] == "true" {
		return nil
	} else {
		this.Flash.Error("您需要登录才能访问该页面。")
		log.Println("非管理员访问了 " + this.Action + " 页面，已自动跳转到登陆页面")
		return this.Redirect(controllers.Admin.SignIn)
	}
}
