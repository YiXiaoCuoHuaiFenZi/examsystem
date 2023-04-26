package controllers

import (
	"examsystem/app/models"
	"log"
	"os"
	"strings"

	"github.com/revel/revel"
)

type Question struct {
	*revel.Controller
}

func (qe Question) Create() revel.Result {
	qe.ViewArgs["adminIDCard"] = qe.Session["adminIDCard"]
	qe.ViewArgs["adminName"] = qe.Session["adminName"]

	return qe.Render()
}

func (qe Question) View() revel.Result {
	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Render()
	}
	defer manager.Close()

	var singleChoiceSummary map[string]int
	singleChoiceSummary, err = manager.GetSingleChoiceSummary()
	if err != nil {
		qe.Flash.Error(err.Error())
		return qe.Render()
	}

	var multipleChoiceSummary map[string]int
	multipleChoiceSummary, err = manager.GetMultipleChoiceSummary()
	if err != nil {
		qe.Flash.Error(err.Error())
		return qe.Render()
	}

	var trueFalseSummary map[string]int
	trueFalseSummary, err = manager.GetTrueFalseSummary()
	if err != nil {
		qe.Flash.Error(err.Error())
		return qe.Render()
	}

	trafficQuestionInfo := make(map[string]int)
	lawQuestionInfo := make(map[string]int)
	constructionQuestionInfo := make(map[string]int)

	if v, ok := singleChoiceSummary["交通"]; ok {
		trafficQuestionInfo["单选"] = v
	} else {
		trafficQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["交通"]; ok {
		trafficQuestionInfo["多选"] = v
	} else {
		trafficQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["交通"]; ok {
		trafficQuestionInfo["判断"] = v
	} else {
		trafficQuestionInfo["判断"] = 0
	}

	if v, ok := singleChoiceSummary["司法"]; ok {
		lawQuestionInfo["单选"] = v
	} else {
		lawQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["司法"]; ok {
		lawQuestionInfo["多选"] = v
	} else {
		lawQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["司法"]; ok {
		lawQuestionInfo["判断"] = v
	} else {
		lawQuestionInfo["判断"] = 0
	}

	if v, ok := singleChoiceSummary["建筑"]; ok {
		constructionQuestionInfo["单选"] = v
	} else {
		constructionQuestionInfo["单选"] = 0
	}

	if v, ok := multipleChoiceSummary["建筑"]; ok {
		constructionQuestionInfo["多选"] = v
	} else {
		constructionQuestionInfo["多选"] = 0
	}

	if v, ok := trueFalseSummary["建筑"]; ok {
		constructionQuestionInfo["判断"] = v
	} else {
		constructionQuestionInfo["判断"] = 0
	}

	results := make(map[string]map[string]int)
	results["交通"] = trafficQuestionInfo
	results["司法"] = lawQuestionInfo
	results["建筑"] = constructionQuestionInfo

	qe.ViewArgs["results"] = results
	qe.ViewArgs["adminIDCard"] = qe.Session["adminIDCard"]
	qe.ViewArgs["adminName"] = qe.Session["adminName"]

	return qe.Render()
}

func (qe Question) PostSingleChoice(singleChoice *models.SingleChoice) revel.Result {
	singleChoice.Type = strings.TrimSpace(singleChoice.Type)
	singleChoice.Description = strings.TrimSpace(singleChoice.Description)
	singleChoice.A = strings.TrimSpace(singleChoice.A)
	singleChoice.B = strings.TrimSpace(singleChoice.B)
	singleChoice.C = strings.TrimSpace(singleChoice.C)
	singleChoice.D = strings.TrimSpace(singleChoice.D)
	singleChoice.Answer = strings.TrimSpace(singleChoice.Answer)

	qe.Validation.Required(singleChoice.Type).Message("请选择试题类别")
	qe.Validation.Required(singleChoice.Description).Message("题目描述不能为空")
	qe.Validation.Required(singleChoice.A).Message("选项A不能为空")
	qe.Validation.Required(singleChoice.B).Message("选项B不能为空")
	qe.Validation.Required(singleChoice.C).Message("选项C不能为空")
	qe.Validation.Required(singleChoice.D).Message("选项D不能为空")
	qe.Validation.Required(singleChoice.Answer).Message("答案不能为空")

	switch singleChoice.Answer {
	case "A":
		singleChoice.Answer = singleChoice.A
	case "B":
		singleChoice.Answer = singleChoice.B
	case "C":
		singleChoice.Answer = singleChoice.C
	case "D":
		singleChoice.Answer = singleChoice.D
	default:
		break
	}

	if qe.Validation.HasErrors() {
		qe.Validation.Keep()
		qe.FlashParams()
		return qe.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Render()
	}
	defer manager.Close()

	err = manager.AddSingleChoice(singleChoice)
	if err != nil {
		qe.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "singleChoice.Description"
		qe.Validation.Errors = append(qe.Validation.Errors, &e)

		qe.Validation.Keep()
		qe.FlashParams()
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}

	qe.Flash.Success("创建题目成功")
	log.Println("创建题目成功：", singleChoice)

	return qe.Redirect(Question.Create)
}

func (qe Question) PostMultipleChoice(multipleChoice *models.MultipleChoice, answers []string) revel.Result {
	multipleChoice.Type = strings.TrimSpace(multipleChoice.Type)
	multipleChoice.Description = strings.TrimSpace(multipleChoice.Description)
	multipleChoice.A = strings.TrimSpace(multipleChoice.A)
	multipleChoice.B = strings.TrimSpace(multipleChoice.B)
	multipleChoice.C = strings.TrimSpace(multipleChoice.C)
	multipleChoice.D = strings.TrimSpace(multipleChoice.D)
	multipleChoice.E = strings.TrimSpace(multipleChoice.E)
	multipleChoice.F = strings.TrimSpace(multipleChoice.F)
	//multipleChoice.Answer = strings.TrimSpace(multipleChoice.Answer)

	qe.Validation.Required(multipleChoice.Type).Message("请选择试题类别")
	qe.Validation.Required(multipleChoice.Description).Message("题目描述不能为空")
	qe.Validation.Required(multipleChoice.A).Message("选项A不能为空")
	qe.Validation.Required(multipleChoice.B).Message("选项B不能为空")
	qe.Validation.Required(multipleChoice.C).Message("选项C不能为空")
	qe.Validation.Required(multipleChoice.D).Message("选项D不能为空")
	qe.Validation.Required(multipleChoice.E).Message("选项E不能为空")
	qe.Validation.Required(multipleChoice.F).Message("选项F不能为空")
	//qe.Validation.Required(len(answers)>0).Message("答案不能为空")
	qe.Validation.Required(answers).Message("答案不能为空")

	var as []string
	for _, a := range answers {
		switch a {
		case "A":
			as = append(as, multipleChoice.A)
		case "B":
			as = append(as, multipleChoice.B)
		case "C":
			as = append(as, multipleChoice.C)
		case "D":
			as = append(as, multipleChoice.D)
		case "E":
			as = append(as, multipleChoice.E)
		case "F":
			as = append(as, multipleChoice.F)
		default:
			break
		}
	}
	multipleChoice.Answer = as

	if qe.Validation.HasErrors() {
		qe.Validation.Keep()
		qe.FlashParams()
		return qe.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}
	defer manager.Close()

	err = manager.AddMultipleChoice(multipleChoice)
	if err != nil {
		qe.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "multipleChoice.Description"
		qe.Validation.Errors = append(qe.Validation.Errors, &e)

		qe.Validation.Keep()
		qe.FlashParams()
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}

	qe.Flash.Success("创建题目成功")
	log.Println("创建题目成功：", multipleChoice)

	return qe.Redirect(Question.Create)
}

func (qe Question) PostTrueFalse(trueFalse *models.TrueFalse) revel.Result {
	trueFalse.Type = strings.TrimSpace(trueFalse.Type)
	trueFalse.Description = strings.TrimSpace(trueFalse.Description)
	trueFalse.Answer = strings.TrimSpace(trueFalse.Answer)

	qe.Validation.Required(trueFalse.Type).Message("请选择试题类别")
	qe.Validation.Required(trueFalse.Description).Message("题目描述不能为空")
	qe.Validation.Required(trueFalse.Answer).Message("答案不能为空")

	if qe.Validation.HasErrors() {
		qe.Validation.Keep()
		qe.FlashParams()
		return qe.Redirect(Question.Create)
	}

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}
	defer manager.Close()

	err = manager.AddTrueFalse(trueFalse)
	if err != nil {
		qe.Validation.Clear()

		var e revel.ValidationError
		e.Message = err.Error()
		e.Key = "trueFalse.Description"
		qe.Validation.Errors = append(qe.Validation.Errors, &e)

		qe.Validation.Keep()
		qe.FlashParams()
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}

	qe.Flash.Success("创建题目成功")
	log.Println("创建题目成功：", trueFalse)

	return qe.Redirect(Question.Create)
}

func (qe Question) PostBatchSingleChoice(batchSingleChoiceFile *os.File, qType string) revel.Result {
	// TODO 文件默认是ascII编码， 需要进行处理
	// 暂时强制要求手动转换为utf8
	scs, err := models.ParseSingleChoiceFile(batchSingleChoiceFile, qType)
	defer batchSingleChoiceFile.Close()

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}
	defer manager.Close()

	successMsg, errorMsg := "", ""
	for _, sc := range scs {
		e := manager.AddSingleChoice(&sc)
		if e != nil {
			m := e.Error() + "：" + sc.Description + "\n"
			log.Println(sc.Description)
			errorMsg += m
			log.Println(m)
		} else {
			successMsg += "创建成功：" + sc.Description + "\n"
			log.Println("创建成功：", sc.Description)
		}
	}

	qe.Flash.Success(successMsg + errorMsg)

	qe.Session["batch"] = "true"
	return qe.Redirect(Question.Create)
}

func (qe Question) PostBatchMultipleChoice(batchMultipleChoiceFile *os.File, qType string) revel.Result {
	// TODO 文件默认是ascII编码， 需要进行处理
	// 暂时强制要求手动转换为utf8
	mcs, err := models.ParseMultipleChoiceFile(batchMultipleChoiceFile, qType)
	defer batchMultipleChoiceFile.Close()

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}
	defer manager.Close()

	successMsg, errorMsg := "", ""
	for _, mc := range mcs {
		err = manager.AddMultipleChoice(&mc)
		if err != nil {
			m := err.Error() + "：" + mc.Description + "\n"
			errorMsg += m
			log.Println(m)
		} else {
			successMsg += "创建成功：" + mc.Description + "\n"
			log.Println("创建成功：", mc)
		}
	}

	qe.Flash.Success(successMsg + errorMsg)

	qe.Session["batch"] = "true"
	return qe.Redirect(Question.Create)
}

func (qe Question) PostBatchTrueFalse(batchTrueFalseFile *os.File, qType string) revel.Result {
	// TODO 文件默认是ascII编码， 需要进行处理
	// 暂时强制要求手动转换为utf8
	tfs, err := models.ParseTrueFalseFile(batchTrueFalseFile, qType)
	defer batchTrueFalseFile.Close()

	manager, err := models.NewDBManager()
	if err != nil {
		qe.Response.Status = 500
		qe.Flash.Error(err.Error())
		return qe.Redirect(Question.Create)
	}
	defer manager.Close()

	successMsg, errorMsg := "", ""
	for _, tf := range tfs {
		err = manager.AddTrueFalse(&tf)
		if err != nil {
			m := err.Error() + "：" + tf.Description + "\n"
			errorMsg += m
			log.Println(m)
		} else {
			successMsg += "创建成功：" + tf.Description + "\n"
			log.Println("创建成功：", tf)
		}
	}

	qe.Flash.Success(successMsg + errorMsg)

	qe.Session["batch"] = "true"
	return qe.Redirect(Question.Create)
}
