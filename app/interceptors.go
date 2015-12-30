package app

import (
	//"ExamSystem/app/controllers" 为什么不对？
	"examsystem/app/controllers"
	"log"

	"github.com/revel/revel"
)

func hasAuthority(this *revel.Controller) revel.Result {
	if this.Session["administrator"] != "true" {
		this.Flash.Error("你没有访问该页面的权限，已自动跳转到首页。")
		log.Println("考生访问了禁止的页面，已自动跳转到考生首页")
		return this.Redirect(controllers.Examinee.Index)
	}
	return nil
}
