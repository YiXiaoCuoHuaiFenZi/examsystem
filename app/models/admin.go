package models

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (manager *DBManager) CreateSuperAdministrator(IDCard string, password string) error {
	t := manager.GetAdminCollection()

	log.Println("检测超级管理员配置")
	count, err := t.CountDocuments(context.TODO(), bson.M{"idcard": IDCard})
	if count != 0 {
		log.Println("超级管理员账号已存在。")
		return nil
	}

	log.Println("超级管理员账号不存在，创建中...")

	var a Admin
	a.IDCard = IDCard
	a.Gender = "天神"
	a.Name = "Administrator"
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	a.Password = p

	if err != nil {
		return err
	}

	_, err = t.InsertOne(context.TODO(), a)
	if err != nil {
		err = errors.New("创建管理员失败")
	} else {
		log.Println("成功创建超级管理员账号.")
	}

	return err
}
func (manager *DBManager) AdminSignUp(signUpAdmin *SignUpAdmin) error {
	t := manager.GetAdminCollection()

	count, err := t.CountDocuments(context.TODO(), bson.M{"idcard": signUpAdmin.IDCard})
	if count != 0 {
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

	_, err = t.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println("创建管理员失败：")
		log.Println(a)
	}

	return err
}

func (manager *DBManager) AdminSignIn(signInAdmin *SignInAdmin) (admin *Admin, err error) {
	t := manager.GetAdminCollection()
	log.Println("AdminSignIn  ", signInAdmin)
	count, err := t.CountDocuments(context.TODO(), bson.M{"idcard": signInAdmin.IDCard})

	if count == 0 {
		log.Println("该管理员不存在")
		err = errors.New("该管理员不存在")
		return
	}

	err = t.FindOne(context.TODO(), bson.M{"idcard": signInAdmin.IDCard}).Decode(&admin)
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
