package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo   Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3      AwsS3      `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

// GetConfigInfo 获取配置信息
func GetConfigInfo() CommonConfig {
	return config
}

type CommonConfig struct {
	StaffGender       []string //员工性别
	StaffIsLeader     []string //员工是否领导
	IsMain            []string //是否主部门
	Nation            []string //民族
	HouseholdType     []string //户籍类型
	Marriage          []string //婚否
	PoliticalOutlook  []string //政治面貌
	Relationship      []string //联系人关系
	Education         []string //学历
	StaffJobType      []string //员工类型
	StaffJobStatus    []string //员工状态
	StaffJobTryPeriod []string //员工试用期
	ExpenseAccount    []string //费用科目
	AgreementType     []string //合同类型
	FileType          []string //合同类型
	SocialType        []string //社保公积金类型
	CredentialType    []string //证件类型
	SalaryType        []string //工资类型
	IsRequired        []string //是否必选款项
}

var config = CommonConfig{
	StaffGender: []string{
		0: "未知",
		1: "男",
		2: "女",
	},
	StaffIsLeader: []string{
		0: "否",
		1: "是",
	},
	IsMain: []string{
		0: "否",
		1: "是",
	},
	Nation: []string{
		1:  "汉族",
		2:  "蒙古族",
		3:  "回族",
		4:  "藏族",
		5:  "维吾尔族",
		6:  "苗族",
		7:  "彝族",
		8:  "壮族",
		9:  "布依族",
		10: "朝鲜族",
		11: "满族",
		12: "侗族",
		13: "瑶族",
		14: "白族",
		15: "土家族",
		16: "哈尼族",
		17: "哈萨克族",
		18: "傣族",
		19: "黎族",
		20: "傈僳族",
		21: "佤族",
		22: "畲族",
		23: "高山族",
		24: "拉祜族",
		25: "水族",
		26: "东乡族",
		27: "纳西族",
		28: "景颇族",
		29: "柯尔克孜族",
		30: "土族",
		31: "达斡尔族",
		32: "仫佬族",
		33: "羌族",
		34: "布朗族",
		35: "撒拉族",
		36: "毛南族",
		37: "仡佬族",
		38: "锡伯族",
		39: "阿昌族",
		40: "普米族",
		41: "塔吉克族",
		42: "怒族",
		43: "乌孜别克族",
		44: "俄罗斯族",
		45: "鄂温克族",
		46: "崩龙族",
		47: "保安族",
		48: "裕固族",
		49: "京族",
		50: "塔塔尔族",
		51: "独龙族",
		52: "鄂伦春族",
		53: "赫哲族",
		54: "门巴族",
		55: "珞巴族",
		56: "基诺族",
	},
	HouseholdType: []string{
		1: "本地城镇",
		2: "本地农村",
		3: "外地城镇[省内]",
		4: "外地农村[省内]",
		5: "外地城镇[省外]",
		6: "外地农村[省外]",
	},
	Marriage: []string{
		1: "已婚",
		2: "未婚",
		3: "其他",
	},
	PoliticalOutlook: []string{
		0: "其他",
		1: "团员",
		2: "党员",
		3: "群众",
	},
	Relationship: []string{
		0: "其他",
		1: "父母",
		2: "配偶",
		3: "子女",
	},
	Education: []string{
		1: "大专以下",
		2: "大专",
		3: "专升本/自考本科",
		4: "本科",
		5: "重点本科",
		6: "硕士",
		7: "重点硕士",
		8: "博士",
	},
	StaffJobType: []string{
		0: "其他",
		1: "全职",
		2: "兼职",
		3: "跟岗实习",
		4: "顶岗实习",
		5: "退休返聘",
		6: "劳务外包",
	},
	StaffJobStatus: []string{
		0: "待入职",
		1: "试用",
		2: "正式",
		3: "待离职",
		4: "已离职",
	},
	StaffJobTryPeriod: []string{
		0: "其他",
		1: "无试用期",
		2: "2个月",
		3: "6个月",
	},
	ExpenseAccount: []string{
		1: "管理费用",
		2: "研发费用",
		3: "生产费用",
		4: "销售费用",
	},
	AgreementType: []string{
		1: "固定期限劳动合同",
		2: "无固定期限劳动合同",
		3: "实习协议",
		4: "外包协议",
		5: "劳务派遣合同",
		6: "返聘协议",
		7: "培训协议",
	},
	FileType: []string{
		0:  "其它",
		1:  "合同附件",
		2:  "身份证(人像)",
		3:  "身份证(国徽)",
		4:  "学历证书",
		5:  "学位证书",
		6:  "前公司离职证明",
		7:  "员工入职申请表",
		8:  "试用期管理规定",
		9:  "个人简历",
		10: "职称/技能证书",
	},
	SocialType: []string{
		1: "深圳社保",
		2: "深圳公积金",
		3: "东莞社保",
		4: "东莞公积金",
	},
	CredentialType: []string{
		0: "其它",
		1: "居民身份证",
		2: "港澳通行证",
	},
	SalaryType: []string{
		1: "基本工资",
		2: "集团经营绩效奖励",
		3: "节日金",
		4: "半年奖",
		5: "年度奖金",
		6: "总部职能体系月度奖金",
		7: "总部金纳斯市场体系月度奖金",
		8: "总部调理中心体系月度奖金",
	},
	IsRequired: []string{
		0: "否",
		1: "是",
	},
}
