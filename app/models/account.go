package models

import (
	"errors"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) SignUp(signUpUser *SignUpUser) error {
	t := this.session.DB(DBName).C(ExamineeCollection)

	i, _ := t.Find(bson.M{"idcard": signUpUser.IDCard}).Count()
	if i != 0 {
		return errors.New("该身份证已被注册")
	}

	var u User
	u.IDCard = signUpUser.IDCard
	u.Gender = signUpUser.Gender
	u.Name = signUpUser.Name
	p, err := bcrypt.GenerateFromPassword([]byte(signUpUser.Password), bcrypt.DefaultCost)
	u.Password = p

	if err != nil {
		return err
	}

	err = t.Insert(u)
	if err != nil {
		log.Println("创建用户失败：")
		log.Println(u)
	}

	return err
}

func (this *DBManager) SignIn(signInUser *SignInUser) (user *User, err error) {
	t := this.session.DB(DBName).C(ExamineeCollection)

	i, _ := t.Find(bson.M{"idcard": signInUser.IDCard}).Count()
	if i == 0 {
		log.Println("该用户不存在")
		err = errors.New("该用户不存在")
		return
	}

	t.Find(bson.M{"idcard": signInUser.IDCard}).One(&user)

	if user.Password == nil {
		log.Println("获取密码错误")
		err = errors.New("获取密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(signInUser.Password))
	if err != nil {
		log.Println("密码不正确")
		err = errors.New("密码不正确")
	}
	return
}
