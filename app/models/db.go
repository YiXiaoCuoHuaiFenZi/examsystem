package models

import (
	"context"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	client *mongo.Client
}

func NewDBManager() (*DBManager, error) {
	revel.Config.SetSection("db")
	uri, found := revel.Config.String("uri")
	if !found {
		log.Fatal("Cannot load database uri from app.conf.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}

	return &DBManager{client}, nil
}

func (manager *DBManager) DataBase(name string) *mongo.Database {
	return manager.client.Database(name)
}

func (manager *DBManager) Collection(name string) *mongo.Collection {
	return manager.client.Database(DBName).Collection(name)
}

func (manager *DBManager) GetExamineeCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(ExamineeCollection)
}

func (manager *DBManager) GetAdminCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(AdminCollection)
}

func (manager *DBManager) GetSingleChoiceCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(SingleChoiceCollection)
}

func (manager *DBManager) GetMultipleChoiceCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(MultipleChoiceCollection)
}

func (manager *DBManager) GetTrueFalseCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(TrueFalseCollection)
}

func (manager *DBManager) GetExamPaperCollection() *mongo.Collection {
	return manager.client.Database(DBName).Collection(ExamPaperCollection)
}

func (manager *DBManager) Close() {
	if err := manager.client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Disconnect database fail!")
	}
}
