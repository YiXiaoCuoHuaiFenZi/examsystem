package models

import (
	"errors"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) SignUp(mu *MockUser) error {
	t := this.session.DB(DBName).C(UserCollection)

	i, _ := t.Find(bson.M{"idcard": mu.IDCard}).Count()
	if i != 0 {
		return errors.New("该身份证已被注册")
	}

	var u User
	u.IDCard = mu.IDCard
	u.Gender = mu.Gender
	u.Name = mu.Name
	p, err := bcrypt.GenerateFromPassword([]byte(mu.Password), bcrypt.DefaultCost)
	u.Password = p

	if err != nil {
		return err
	}

	err = t.Insert(u)
	if err == nil {
		log.Println("成功创建用户：")
		log.Println(u)
	}

	return err
}

//func (manager *DBManager) LoginUser(lu *LoginUser) (user *User, err error) {
//	uc := manager.session.DB(DbName).C(UserCollection)

//	i, _ := uc.Find(bson.M{"nickname": lu.NickName}).Count()
//	if i == 0 {
//		fmt.Println("该用户不存在")
//		err = errors.New("该用户不存在")
//		return
//	}

//	uc.Find(bson.M{"nickname": lu.NickName}).One(&user)

//	if user.Password == nil {
//		err = errors.New("获取密码错误")
//		return
//	}

//	err = bcrypt.CompareHashAndPassword(user.Password, []byte(lu.Password))
//	if err != nil {
//		fmt.Println("密码不正确")
//		err = errors.New("密码不正确")
//	}
//	return
//}
