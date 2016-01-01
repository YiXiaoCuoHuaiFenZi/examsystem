package app

import (
	//"ExamSystem/app/controllers" 为什么不对？
	"examsystem/app/controllers"
	"log"

	"github.com/revel/revel"
)

func beforeAdminController(this *revel.Controller) revel.Result {
	// 设置游客的访问权限
	if this.Session["administrator"] == "" && this.Session["examinee"] == "" {
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
			log.Println("管理员 " + this.Session["adminName"] + " 访问了 " + this.Action + " 页面")
			return nil
		}
	}

	return this.Redirect(controllers.Admin.SignIn)
}

func beforeExamineeController(this *revel.Controller) revel.Result {
	// 设置游客的访问权限
	if this.Session["administrator"] == "" && this.Session["examinee"] == "" {
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
			this.Action == "Examinee.Index" || this.Action == "Examinee.SignOut" {
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
			log.Println("管理员 " + this.Session["adminName"] + " 访问了 " + this.Action + " 页面")
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
