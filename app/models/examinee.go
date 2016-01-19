package models

import (
	"errors"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) SignUp(signUpExaminee *SignUpExaminee) error {
	t := this.session.DB(DBName).C(ExamineeCollection)

	i, _ := t.Find(bson.M{"idcard": signUpExaminee.IDCard}).Count()
	if i != 0 {
		return errors.New("该身份证已被注册")
	}

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

	err = t.Insert(e)
	if err != nil {
		log.Println("创建考生失败：")
		log.Println(e)
	}

	return err
}

func (this *DBManager) SignIn(signInExaminee *SignInExaminee) (examinee *Examinee, err error) {
	t := this.session.DB(DBName).C(ExamineeCollection)

	i, _ := t.Find(bson.M{"idcard": signInExaminee.IDCard}).Count()
	if i == 0 {
		log.Println("该用户不存在")
		err = errors.New("该用户不存在")
		return
	}

	t.Find(bson.M{"idcard": signInExaminee.IDCard}).One(&examinee)

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

func (this *DBManager) GetAllExaminee() ([]Examinee, error) {
	t := this.session.DB(DBName).C(ExamineeCollection)

	count, err := t.Count()
	log.Println("共有考生 ", count, "位")

	examinees := []Examinee{}
	err = t.Find(nil).All(&examinees)

	return examinees, err
}

func (this *DBManager) GetExamineeByIDCard(idCard string) (Examinee, error) {
	t := this.session.DB(DBName).C(ExamineeCollection)

	examinee := Examinee{}
	err := t.Find(bson.M{"idcard": idCard}).One(&examinee)

	return examinee, err
}

func (this *DBManager) UpdateExaminee(examinee *Examinee) error {
	t := this.session.DB(DBName).C(ExamineeCollection)

	var oldExp Examinee
	err := t.Find(bson.M{"idcard": examinee.IDCard}).One(&oldExp)
	if err != nil {
		return err
	}

	tempInfo := oldExp
	tempInfo.ExamType = examinee.ExamPaper.Type
	tempInfo.ExamStatus = examinee.ExamStatus
	tempInfo.ExamPaper = examinee.ExamPaper
	tempInfo.Score = examinee.Score

	err = t.Update(oldExp, tempInfo)
	if err != nil {
		log.Println("更新考生信息失败：")
		return err
	}
	return nil
}
