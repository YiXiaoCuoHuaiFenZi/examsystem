package app

import (
	"examsystem/app/controllers"

	_ "github.com/revel/modules"
	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	// 注册模板里的字符串相加函数
	revel.TemplateFuncs["addOne"] = func(a int) int {
		return a + 1
	}

	revel.TemplateFuncs["lessThan"] = func(a, b int) bool {
		return a < b
	}

	revel.TemplateFuncs["eq"] = func(a, b interface{}) bool {
		return a == b
	}

	revel.TemplateFuncs["greaterThan"] = func(a, b int) bool {
		return a > b
	}

	revel.TemplateFuncs["len"] = func(a []interface{}) int {
		return len(a)
	}

	revel.TemplateFuncs["has"] = func(a []string, b string) bool {
		r := false
		for _, v := range a {
			if v == b {
				r = true
			}
		}

		return r
	}

	//revel.InterceptFunc(checkUser, revel.BEFORE, &App{})
	revel.InterceptFunc(beforeAdminController, revel.BEFORE, &controllers.Admin{})
	revel.InterceptFunc(beforeExamineeController, revel.BEFORE, &controllers.Examinee{})
	//revel.InterceptFunc(beforeAboutController, revel.BEFORE, &controllers.About{})
	revel.InterceptFunc(beforeExamPaperController, revel.BEFORE, &controllers.ExamPaper{})
	revel.InterceptFunc(beforeQuestionController, revel.BEFORE, &controllers.Question{})
}

// HeaderFilter TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
