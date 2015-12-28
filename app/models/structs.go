package models

// 考生信息数据结构--用于注册
type SignUpUser struct {
	IDCard          string // 身份证号
	Name            string // 考生姓名
	Gender          string // 考生性别
	Password        string // 账号密码
	ConfirmPassword string // 确认密码
}

// 考生信息数据结构--用于登录
type SignInUser struct {
	IDCard   string // 身份证号
	Name     string // 考生姓名
	Gender   string // 考生性别
	Password string // 账号密码
}

// 考生信息数据结构--真实数据
type User struct {
	IDCard   string // 身份证号
	Name     string // 考生姓名
	Gender   string // 考生性别
	Password []byte // 账号密码
}

// 管理员信息数据结构--用于注册
type SignUpAdmin struct {
	IDCard          string // 身份证号
	Name            string // 管理员姓名
	Gender          string // 管理员性别
	Password        string // 账号密码
	ConfirmPassword string // 确认密码
}

// 管理员信息数据结构--用于登录
type SignInAdmin struct {
	IDCard   string // 身份证号
	Name     string // 管理员姓名
	Gender   string // 管理员性别
	Password string // 账号密码
}

// 管理员信息数据结构--真实数据
type Admin struct {
	IDCard string // 身份证号
	Name   string // 管理员姓名
	Gender string // 管理员性别
	Level  int    // 管理员权限级别
}
