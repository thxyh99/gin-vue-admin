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
	StaffStatus       []string //员工状态
	StaffInfoType     []string //员工考勤类型
	StaffOfficeRank   []string //员工职级(内勤)
	StaffMarketRank   []string //员工职级(市场)
	Nation            []string //民族
	HouseholdType     []string //户籍类型
	Marriage          []string //婚否
	PoliticalOutlook  []string //政治面貌
	Relationship      []string //联系人关系
	Education         []string //学历
	StaffJobType      []string //员工工作类型
	StaffJobStatus    []string //员工状态
	StaffJobTryPeriod []string //员工试用期
}

var config = CommonConfig{
	StaffGender: []string{
		0: "未知",
		1: "男",
		2: "女",
	},
	StaffStatus: []string{
		1: "已激活",
		2: "已禁用",
		4: "未激活",
		5: "退出企业",
	},
	StaffInfoType: []string{
		1: "内勤",
		2: "市场",
	},
	StaffOfficeRank: []string{
		1:  "实习生",
		2:  "职员5级",
		3:  "职员4级",
		4:  "职员3级",
		5:  "职员2级",
		6:  "职员1级",
		7:  "专员3级",
		8:  "专员2级",
		9:  "专员1级",
		10: "主管3级",
		11: "主管2级",
		12: "主管1级",
		13: "资深专员3级",
		14: "资深专员2级",
		15: "资深专员1级",
		16: "资深主管3级",
		17: "资深主管2级",
		18: "资深主管1级",
		19: "专业副经理3级",
		20: "专业副经理2级",
		21: "专业副经理1级",
		22: "专业经理3级",
		23: "专业经理2级",
		24: "专业经理1级",
		25: "副经理3级",
		26: "副经理2级",
		27: "副经理1级",
		28: "经理3级",
		29: "经理2级",
		30: "经理1级",
		31: "副监3级",
		32: "副监2级",
		33: "副监1级",
		34: "总监3级",
		35: "总监2级",
		36: "总监1级",
		37: "副总3级",
		38: "副总2级",
		39: "副总1级",
		40: "总经理3级",
		41: "总经理2级",
		42: "总经理1级",
	},
	StaffMarketRank: []string{},
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
		0: "其他",
		1: "小学",
		2: "初中",
		3: "高中",
		4: "中专",
		5: "大专",
		6: "本科",
		7: "硕士",
		8: "博士",
	},
	StaffJobType: []string{
		1: "全职",
		2: "兼职",
		3: "实习",
		4: "劳务外包",
		5: "无类型",
	},
	StaffJobStatus: []string{
		0: "无状态",
		1: "试用",
		2: "正式",
		3: "待离职",
	},
	StaffJobTryPeriod: []string{
		0: "其他",
		1: "无试用期",
		2: "1个月",
		3: "2个月",
		4: "3个月",
		5: "4个月",
		6: "5个月",
		7: "6个月",
	},
}
