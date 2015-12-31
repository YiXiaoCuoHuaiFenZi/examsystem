package app

import (
	//"ExamSystem/app/controllers" 为什么不对？
	"examsystem/app/controllers"
	"log"

	"github.com/revel/revel"
)

func checkAuthority(this *revel.Controller) revel.Result {
	// 没有登陆的用户、游客访问了管理员后台页面
	// 自动跳转到管理员登陆页面
	if this.Session["administrator"] == "" && this.Session["examinee"] == "" {
		if this.Action == "Admin.SignIn" || this.Action == "Admin.PostSignIn" {
			return nil
		} else {
			this.Flash.Error("您访问了需要管理员权限才能访问的页面，请先登陆，已为您跳转到管理员登陆页面。")
			log.Println("未登录用户访问了管理员后台页面，已自动跳转到管理员登陆页面")
			return this.Redirect(controllers.Admin.SignIn)
		}
	}

	// 已经登陆的考生访问了管理员后台页面
	// 自动跳转到考生首页
	if this.Session["examinee"] == "true" {
		this.Flash.Error("您没有访问该页面的权限，已自动跳转到首页。")
		log.Println("考生访问了禁止的页面，已自动跳转到考生首页")
		return this.Redirect(controllers.Examinee.Index)
	}

	// 已经登陆过的管理员访问页面
	// 打印log日志
	if this.Session["administrator"] == "true" {
		log.Println("管理员 " + this.Session["adminName"] + " 访问了 " + this.Name + " 页面")
		return nil
	}

	return nil
}
