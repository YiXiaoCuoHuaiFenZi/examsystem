package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) AdminSignUp(signUpAdmin *SignUpAdmin) error {
	t := this.session.DB(DBName).C(AdminCollection)

	i, _ := t.Find(bson.M{"idcard": signUpAdmin.IDCard}).Count()
	if i != 0 {
		return errors.New("该身份证已被注册")
	}

	var a Admin
	a.IDCard = signUpAdmin.IDCard
	a.Gender = signUpAdmin.Gender
	a.Name = signUpAdmin.Name
	p, err := bcrypt.GenerateFromPassword([]byte(signUpAdmin.Password), bcrypt.DefaultCost)
	a.Password = p

	if err != nil {
		return err
	}

	err = t.Insert(a)
	if err != nil {
		log.Println("创建管理员失败：")
		log.Println(a)
	}

	return err
}

func (this *DBManager) AdminSignIn(signInAdmin *SignInAdmin) (admin *Admin, err error) {
	t := this.session.DB(DBName).C(AdminCollection)

	i, _ := t.Find(bson.M{"idcard": signInAdmin.IDCard}).Count()
	if i == 0 {
		log.Println("该管理员不存在")
		err = errors.New("该管理员不存在")
		return
	}

	t.Find(bson.M{"idcard": signInAdmin.IDCard}).One(&admin)

	if admin.Password == nil {
		log.Println("获取密码错误")
		err = errors.New("获取密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword(admin.Password, []byte(signInAdmin.Password))
	if err != nil {
		log.Println("密码不正确")
		err = errors.New("密码不正确")
	}
	return
}
