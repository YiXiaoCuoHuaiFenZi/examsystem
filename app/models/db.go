package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

const (
	DBName                   = "ExamSystem"
	ExamineeCollection       = "examinee"
	AdminCollection          = "admin"
	SingleChoiceCollection   = "singlechoice"
	MultipleChoiceCollection = "multiplechoice"
	TrueFalseCollection      = "truefalse"
	ExamPaperCollection      = "exampaper"
)

type DBManager struct {
	session *mgo.Session
}

func NewDBManager() (*DBManager, error) {
	revel.Config.SetSection("db")
	ip, found := revel.Config.String("ip")
	if !found {
		revel.ERROR.Fatal("Cannot load database ip from app.conf")
	}

	session, err := mgo.Dial(ip)
	if err != nil {
		return nil, err
	}

	return &DBManager{session}, nil
}

func (manager *DBManager) DataBase(name string) *mgo.Database {
	return manager.session.DB(name)
}

func (manager *DBManager) Collection(name string) *mgo.Collection {
	return manager.session.DB(DBName).C(name)
}

func (manager *DBManager) Close() {
	manager.session.Close()
}
