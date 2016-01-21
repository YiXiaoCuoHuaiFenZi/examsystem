package models

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
	IDCard   string // 身份证号
	Name     string // 管理员姓名
	Gender   string // 管理员性别
	Password []byte // 账号密码
	Level    int    // 管理员权限级别
}

// 考生信息数据结构--用于注册
type SignUpExaminee struct {
	IDCard          string // 身份证号
	Name            string // 考生姓名
	Gender          string // 考生性别
	Mobile          string // 手机号码
	Password        string // 账号密码
	ConfirmPassword string // 确认密码
}

// 考生信息数据结构--用于登录
type SignInExaminee struct {
	IDCard   string // 身份证号
	Name     string // 考生姓名
	Gender   string // 考生性别
	Password string // 账号密码
}

// 考生信息数据结构--真实数据
type Examinee struct {
	IDCard    string    // 身份证号
	Name      string    // 考生姓名
	Gender    string    // 考生性别
	Password  []byte    // 账号密码
	Mobile    string    // 手机号码
	ExamPaper ExamPaper // 所有试卷
}

type SingleChoice struct {
	Type         string // 所属考试类别
	Discription  string // 题目描述
	A            string // 选项A
	B            string // 选项B
	C            string // 选项C
	D            string // 选项D
	Answer       string // 正确答案
	ActualAnswer string // 考生实际答案
}

type MultipleChoice struct {
	Type         string   // 所属考试类别
	Discription  string   // 题目描述
	A            string   // 选项A
	B            string   // 选项B
	C            string   // 选项C
	D            string   // 选项D
	E            string   // 选项E
	F            string   // 选项F
	Answer       []string // 正确答案
	ActualAnswer []string // 考生实际答案
}

type TrueFalse struct {
	Type         string // 所属考试类别
	Discription  string // 题目描述
	Answer       string // 正确答案
	ActualAnswer string // 考生实际答案
}

type SCWithPage struct {
	Page int          // 考试时所在试卷上的页码标识
	SC   SingleChoice // 单选题信息
}

type MCWithPage struct {
	Page int            // 考试时所在试卷上的页码标识
	MC   MultipleChoice // 多选题信息
}

type TFWithPage struct {
	Page int       // 考试时所在试卷上的页码标识
	TF   TrueFalse // 判断题信息
}

// 考试状态
const (
	NoExam     = "无考试" // 无考试
	UnFinished = "待考"  // 有未完成的考试
	Doing      = "考试中" // 考试进行中
	Done       = "完成"  // 考试完成
)

type ExamPaper struct {
	Type         string           // 所属考试类别
	Title        string           // 试卷标题
	Discription  string           // 试卷描述
	Score        float32          // 总分值
	ActualScore  float32          // 考生实际得分
	Time         int              // 考试时间
	TimeStamp    string           // 时间戳
	IDCode       string           // 编号
	CreateMethod string           // 试卷的生成方式：随机、套题
	Status       string           // 考试状态
	SCCount      int              // 单选题数量
	SCScore      float32          // 单选题每题分值
	SC           []SingleChoice   // 单选题
	MCCount      int              // 多选题数量
	MCScore      float32          // 多选题每题分值
	MC           []MultipleChoice // 多选题
	TFCount      int              // 判断题数量
	TFScore      float32          // 判断题每题分值
	TF           []TrueFalse      // 判断题
}
