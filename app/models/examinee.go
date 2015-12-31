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