package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

const (
	DBName         = "ExamSystem"
	UserCollection = "user"
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

func (manager *DBManager) Close() {
	manager.session.Close()
}
