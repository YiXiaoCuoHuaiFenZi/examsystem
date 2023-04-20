package models

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (manager *DBManager) GetAllExamPaper() ([]ExamPaper, error) {
	t := manager.GetExamPaperCollection()

	var examPapers []ExamPaper
	cursor, err := t.Find(context.TODO(), bson.M{})
	for cursor.Next(context.TODO()) {
		var examPaper = ExamPaper{}
		err = cursor.Decode(&examPaper)
		if err != nil {
			log.Println(err)
		}
		examPapers = append(examPapers, examPaper)
	}

	if len(examPapers) == 0 {
		log.Println("试卷库没有可用试卷")
		return nil, errors.New("试卷库没有可用试卷")
	}

	return examPapers, nil
}

func (manager *DBManager) AddExamPaper(exp *ExamPaper) error {
	t := manager.GetExamPaperCollection()

	count, err := t.CountDocuments(context.TODO(), bson.M{"title": exp.Title})
	if count != 0 {
		return errors.New("该试卷已经存在")
	}

	_, err = t.InsertOne(context.TODO(), exp)
	if err != nil {
		log.Println("创建试卷失败：", exp)
	}

	return err
}

func (manager *DBManager) GetExamPaperByTitle(title string) (ExamPaper, error) {
	t := manager.GetExamPaperCollection()

	exp := ExamPaper{}
	err := t.FindOne(context.TODO(), bson.M{"title": title}).Decode(&exp)
	if err != nil {
		return exp, err
	}

	return exp, err
}
