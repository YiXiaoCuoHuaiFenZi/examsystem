package models

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) GetAllExamPaper() ([]ExamPaper, error) {
	t := this.session.DB(DBName).C(ExamPaperCollection)

	exp := []ExamPaper{}

	err := t.Find(nil).All(&exp)
	if err != nil {
		return nil, err
	}

	if len(exp) == 0 {
		log.Println("试卷库没有可用试卷")
		return nil, errors.New("试卷库没有可用试卷")
	}

	return exp, nil
}

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

func (this *DBManager) GetExamPaperByTitle(title string) (ExamPaper, error) {
	t := this.session.DB(DBName).C(ExamPaperCollection)

	exp := ExamPaper{}
	err := t.Find(bson.M{"title": title}).One(&exp)
	if err != nil {
		return exp, err
	}

	return exp, err
}
