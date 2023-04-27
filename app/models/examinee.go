package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (manager *DBManager) SignUp(signUpExaminee *SignUpExaminee) error {
	t := manager.GetExamineeCollection()

	var examinee Examinee
	err := t.FindOne(context.TODO(), bson.M{"idcard": signUpExaminee.IDCard}).Decode(&examinee)
	if examinee.IDCard != "" {
		return errors.New("该身份证已被注册")
	}

	if err == mongo.ErrNoDocuments {
		var e Examinee
		e.IDCard = signUpExaminee.IDCard
		e.Gender = signUpExaminee.Gender
		e.Name = signUpExaminee.Name
		e.Mobile = signUpExaminee.Mobile
		p, err := bcrypt.GenerateFromPassword([]byte(signUpExaminee.Password), bcrypt.DefaultCost)
		e.Password = p

		if err != nil {
			return err
		}

		_, err = t.InsertOne(context.TODO(), e)
		if err != nil {
			log.Println("创建考生失败：")
			log.Println(e)
			return errors.New("创建考生失败")
		}
		return nil
	}

	return err
}

func (manager *DBManager) SignIn(signInExaminee *SignInExaminee) (examinee *Examinee, err error) {
	t := manager.GetExamineeCollection()

	err = t.FindOne(context.TODO(), bson.M{"idcard": signInExaminee.IDCard}).Decode(&examinee)
	if err == mongo.ErrNoDocuments {
		log.Println("考生不存在")
		err = errors.New("考生不存在")
		return
	}

	if examinee.Password == nil {
		log.Println("获取密码错误")
		err = errors.New("获取密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword(examinee.Password, []byte(signInExaminee.Password))
	if err != nil {
		log.Println("密码不正确")
		err = errors.New("密码不正确")
	}
	return
}

func (manager *DBManager) GetAllExaminee() ([]Examinee, error) {
	t := manager.GetExamineeCollection()

	count, err := t.CountDocuments(context.TODO(), bson.D{})
	log.Println("共有考生 ", count, "位")
	if err != nil {
		log.Println(err)
		err = errors.New("没有查到考生数据")
		return nil, err
	}

	var examinees []Examinee
	cursor, err := t.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var examinee = Examinee{}
		err = cursor.Decode(&examinee)
		if err != nil {
			log.Println(err)
		}
		examinees = append(examinees, examinee)
	}

	return examinees, err
}

func (manager *DBManager) GetExamineeByIDCard(idCard string) (Examinee, error) {
	t := manager.GetExamineeCollection()

	examinee := Examinee{}
	err := t.FindOne(context.TODO(), bson.M{"idcard": idCard}).Decode(&examinee)
	if err == mongo.ErrNoDocuments {
		log.Println("考生不存在")
		err = errors.New("考生不存在")
	}

	return examinee, err
}

func (manager *DBManager) UpdateExamPaper(idCard string, examPaper ExamPaper) error {
	t := manager.GetExamineeCollection()

	var oldExp Examinee
	err := t.FindOne(context.TODO(), bson.M{"idcard": idCard}).Decode(&oldExp)
	if err == mongo.ErrNoDocuments {
		log.Println("考生不存在")
		err = errors.New("考生不存在")
		return err
	}

	filter := bson.D{{"idcard", idCard}}
	update := bson.D{{"$set", bson.D{{"exampaper", examPaper}}}}
	result, err := t.UpdateOne(context.TODO(), filter, update)
	log.Println("UpdateExaminee：", result)
	if err != nil {
		log.Println("更新考生信息失败：")
		return err
	}

	return nil
}
