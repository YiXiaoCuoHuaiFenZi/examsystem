package models

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) AddExamPaper(exp *ExamPaper) error {
	t := this.session.DB(DBName).C(ExamPaperCollection)

	i, _ := t.Find(bson.M{"title": exp.Title}).Count()
	if i != 0 {
		return errors.New("该试卷已经存在")
	}

	err := t.Insert(exp)
	if err != nil {
		log.Println("创建试卷失败：", exp)
	}

	return err
}
