package models

// 考生信息数据结构
type MockUser struct {
	IDCard          string  // 身份证号
	Name            string  // 考生姓名	
	Gender          string  // 考生性别
	Password        string  // 账号密码
	ConfirmPassword string  // 确认密码 
}

// 考生信息数据结构
type User struct {
	IDCard          string  // 身份证号
	Name            string  // 考生姓名	
	Gender          string  // 考生性别
	Password        string  // 账号密码
}