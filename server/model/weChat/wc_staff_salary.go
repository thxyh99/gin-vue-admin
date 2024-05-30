// 自动生成模板WcStaffSalary
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 工资列表 结构体  WcStaffSalary
type WcStaffSalary struct {
	global.GVA_MODEL
	Type                    *int       `json:"type" form:"type" gorm:"column:type;comment:工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金);size:2;"` //工资类型(1:基本工资 2:集团经营绩效奖励 3:节日金 4:半年奖 5:年度奖金 6:总部职能体系月度奖金 7:总部金纳斯市场体系月度奖金 8:总部调理中心体系月度奖金)
	StaffId                 *int       `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`                                                     //员工ID
	DepartmentFirst         string     `json:"departmentFirst" form:"departmentFirst" gorm:"column:department_first;comment:一级部门;size:200;"`                                               //一级部门
	DepartmentSecond        string     `json:"departmentSecond" form:"departmentSecond" gorm:"column:department_second;comment:二级部门;size:200;"`                                            //二级部门
	Name                    string     `json:"name" form:"name" gorm:"column:name;comment:姓名;size:200;" binding:"required"`                                                                //姓名
	EmploymentDate          *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职时间;"`                                                           //入职时间
	Deadline                *time.Time `json:"deadline" form:"deadline" gorm:"column:deadline;comment:截止日期;"`                                                                              //截止日期
	CompanyAge              *int       `json:"companyAge" form:"companyAge" gorm:"column:company_age;comment:司龄;size:3;"`                                                                  //司龄
	SalaryDays              *int       `json:"salaryDays" form:"salaryDays" gorm:"column:salary_days;comment:发薪天数;size:3;"`                                                                //发薪天数
	Overtime                *float64   `json:"overtime" form:"overtime" gorm:"column:overtime;comment:加班情况;size:5;"`                                                                       //加班情况
	Absence                 *float64   `json:"absence" form:"absence" gorm:"column:absence;comment:缺勤情况;size:5;"`                                                                          //缺勤情况
	Expense                 string     `json:"expense" form:"expense" gorm:"column:expense;comment:费用科目;size:100;"`                                                                        //费用科目
	Position                string     `json:"position" form:"position" gorm:"column:position;comment:职位/职级;size:200;"`                                                                    //职位/职级
	Education               string     `json:"education" form:"education" gorm:"column:education;comment:学历;size:100;"`                                                                    //学历
	GradeSalary             *float64   `json:"gradeSalary" form:"gradeSalary" gorm:"column:grade_salary;comment:等级工资;size:10;"`                                                            //等级工资
	SkillSalary             *float64   `json:"skillSalary" form:"skillSalary" gorm:"column:skill_salary;comment:岗位工资/职称技能津贴;size:10;"`                                                     //岗位工资/职称技能津贴
	SecrecySalary           *float64   `json:"secrecySalary" form:"secrecySalary" gorm:"column:secrecy_salary;comment:保密工资;size:10;"`                                                      //保密工资
	CpiSalary               *float64   `json:"cpiSalary" form:"cpiSalary" gorm:"column:cpi_salary;comment:CPI工资;size:10;"`                                                                 //CPI工资
	AllowancesSalary        *float64   `json:"allowancesSalary" form:"allowancesSalary" gorm:"column:allowances_salary;comment:津贴;size:10;"`                                               //津贴
	AgeAllowanceSalary      *float64   `json:"ageAllowanceSalary" form:"ageAllowanceSalary" gorm:"column:age_allowance_salary;comment:司龄津贴;size:10;"`                                      //司龄津贴
	EducationSalary         *float64   `json:"educationSalary" form:"educationSalary" gorm:"column:education_salary;comment:学历工资;size:10;"`                                                //学历工资
	FilialSalary            *float64   `json:"filialSalary" form:"filialSalary" gorm:"column:filial_salary;comment:尽孝基金;size:10;"`                                                         //尽孝基金
	PerformanceSalary       *float64   `json:"performanceSalary" form:"performanceSalary" gorm:"column:performance_salary;comment:业绩奖;size:10;"`                                           //业绩奖
	FestivalSalary          *float64   `json:"festivalSalary" form:"festivalSalary" gorm:"column:festival_salary;comment:节日金;size:10;"`                                                    //节日金
	BirthdaySalary          *float64   `json:"birthdaySalary" form:"birthdaySalary" gorm:"column:birthday_salary;comment:生日金;size:10;"`                                                    //生日金
	HighTemperatureSubsidy  *float64   `json:"highTemperatureSubsidy" form:"highTemperatureSubsidy" gorm:"column:high_temperature_subsidy;comment:高温补贴;size:10;"`                          //高温补贴
	FullSalary              *float64   `json:"fullSalary" form:"fullSalary" gorm:"column:full_salary;comment:全勤;size:10;"`                                                                 //全勤
	LeaderSubsidy           *float64   `json:"leaderSubsidy" form:"leaderSubsidy" gorm:"column:leader_subsidy;comment:组长补助;size:10;"`                                                      //组长补助
	PackagingSubsidy        *float64   `json:"packagingSubsidy" form:"packagingSubsidy" gorm:"column:packaging_subsidy;comment:分包装间;size:10;"`                                             //分包装间
	JobSubsidy              *float64   `json:"jobSubsidy" form:"jobSubsidy" gorm:"column:job_subsidy;comment:岗位补助;size:10;"`                                                               //岗位补助
	DormitorySubsidy        *float64   `json:"dormitorySubsidy" form:"dormitorySubsidy" gorm:"column:dormitory_subsidy;comment:宿舍补助;size:10;"`                                             //宿舍补助
	MealSubsidy             *float64   `json:"mealSubsidy" form:"mealSubsidy" gorm:"column:meal_subsidy;comment:餐补;size:10;"`                                                              //餐补
	TelephoneSubsidy        *float64   `json:"telephoneSubsidy" form:"telephoneSubsidy" gorm:"column:telephone_subsidy;comment:电话补助;size:10;"`                                             //电话补助
	TransportSubsidy        *float64   `json:"transportSubsidy" form:"transportSubsidy" gorm:"column:transport_subsidy;comment:交通补助;size:10;"`                                             //交通补助
	TravelSubsidy           *float64   `json:"travelSubsidy" form:"travelSubsidy" gorm:"column:travel_subsidy;comment:出差补助;size:10;"`                                                      //出差补助
	OtherSubsidy            *float64   `json:"otherSubsidy" form:"otherSubsidy" gorm:"column:other_subsidy;comment:其他(补助);size:10;"`                                                       //其他(补助)
	TotalSubsidy            *float64   `json:"totalSubsidy" form:"totalSubsidy" gorm:"column:total_subsidy;comment:合计(补助);size:10;"`                                                       //合计(补助)
	OvertimeFee             *float64   `json:"overtimeFee" form:"overtimeFee" gorm:"column:overtime_fee;comment:加班费;size:10;"`                                                             //加班费
	AbsenceFee              *float64   `json:"absenceFee" form:"absenceFee" gorm:"column:absence_fee;comment:缺勤;size:10;"`                                                                 //缺勤
	TotalFee                *float64   `json:"totalFee" form:"totalFee" gorm:"column:total_fee;comment:合计(加班及缺勤);size:10;"`                                                                //合计(加班及缺勤)
	ShouldPay               *float64   `json:"shouldPay" form:"shouldPay" gorm:"column:should_pay;comment:应发工资;size:10;"`                                                                  //应发工资
	PensionPay              *float64   `json:"pensionPay" form:"pensionPay" gorm:"column:pension_pay;comment:养老保险;size:10;"`                                                               //养老保险
	MedicalPay              *float64   `json:"medicalPay" form:"medicalPay" gorm:"column:medical_pay;comment:医疗保险;size:10;"`                                                               //医疗保险
	UnemploymentPay         *float64   `json:"unemploymentPay" form:"unemploymentPay" gorm:"column:unemployment_pay;comment:失业保险;size:10;"`                                                //失业保险
	HousingPay              *float64   `json:"housingPay" form:"housingPay" gorm:"column:housing_pay;comment:住房公积金;size:10;"`                                                              //住房公积金
	SpecialTax              *float64   `json:"specialTax" form:"specialTax" gorm:"column:special_tax;comment:专项抵扣;size:10;"`                                                               //专项抵扣
	ThisPayTax              *float64   `json:"thisPayTax" form:"thisPayTax" gorm:"column:this_pay_tax;comment:本次应缴税所得额;size:10;"`                                                          //本次应缴税所得额
	LastTotalPayTax         *float64   `json:"lastTotalPayTax" form:"lastTotalPayTax" gorm:"column:last_total_pay_tax;comment:上次累计应缴预缴所得额;size:10;"`                                       //上次累计应缴预缴所得额
	ThisTotalPayTax         *float64   `json:"thisTotalPayTax" form:"thisTotalPayTax" gorm:"column:this_total_pay_tax;comment:本次累计应缴预缴所得额;size:10;"`                                       //本次累计应缴预缴所得额
	TotalPayTax             *float64   `json:"totalPayTax" form:"totalPayTax" gorm:"column:total_pay_tax;comment:累计税额;size:10;"`                                                           //累计税额
	LastTotalPaidTax        *float64   `json:"lastTotalPaidTax" form:"lastTotalPaidTax" gorm:"column:last_total_paid_tax;comment:上次累计已扣个税;size:10;"`                                       //上次累计已扣个税
	ThisTotalPaidTax        *float64   `json:"thisTotalPaidTax" form:"thisTotalPaidTax" gorm:"column:this_total_paid_tax;comment:本次累计已缴个税;size:10;"`                                       //本次累计已缴个税
	ThisShouldPayTax        *float64   `json:"thisShouldPayTax" form:"thisShouldPayTax" gorm:"column:this_should_pay_tax;comment:本次应扣个税;size:10;"`                                         //本次应扣个税
	PreFilialSalary         *float64   `json:"preFilialSalary" form:"preFilialSalary" gorm:"column:pre_filial_salary;comment:预存尽孝基金;size:10;"`                                             //预存尽孝基金
	Other                   *float64   `json:"other" form:"other" gorm:"column:other;comment:其他;size:10;"`                                                                                 //其他
	Sent                    *float64   `json:"sent" form:"sent" gorm:"column:sent;comment:已发;size:10;"`                                                                                    //已发
	ActualPay               *float64   `json:"actualPay" form:"actualPay" gorm:"column:actual_pay;comment:实发小计;size:10;"`                                                                  //实发小计
	MonthlyCommission       *float64   `json:"monthlyCommission" form:"monthlyCommission" gorm:"column:monthly_commission;comment:月度提成补发;size:10;"`                                        //月度提成补发
	AnnualCommission        *float64   `json:"annualCommission" form:"annualCommission" gorm:"column:annual_commission;comment:年度提成;size:10;"`                                             //年度提成
	SubBrandCommission      *float64   `json:"subBrandCommission" form:"subBrandCommission" gorm:"column:sub_brand_commission;comment:子品牌经营绩效奖励;size:10;"`                                 //子品牌经营绩效奖励
	BonusPoints             *float64   `json:"bonusPoints" form:"bonusPoints" gorm:"column:bonus_points;comment:奖金点;size:10;"`                                                             //奖金点
	AttendanceRatio         *float64   `json:"attendanceRatio" form:"attendanceRatio" gorm:"column:attendance_ratio;comment:出勤系数;size:10;"`                                                //出勤系数
	MonthlyRating           *float64   `json:"monthlyRating" form:"monthlyRating" gorm:"column:monthly_rating;comment:月度评分;size:10;"`                                                      //月度评分
	PersonalCount           *float64   `json:"personalCount" form:"personalCount" gorm:"column:personal_count;comment:个人计点;size:10;"`                                                      //个人计点
	LateDonation            *float64   `json:"lateDonation" form:"lateDonation" gorm:"column:late_donation;comment:迟到乐捐;size:10;"`                                                         //迟到乐捐
	MonthlyBonus            *float64   `json:"monthlyBonus" form:"monthlyBonus" gorm:"column:monthly_bonus;comment:月度奖金;size:10;"`                                                         //月度奖金
	Remark                  string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:200;"`                                                                             //备注
	Region                  string     `json:"region" form:"region" gorm:"column:region;comment:区域;size:100;"`                                                                             //区域
	AnnualTask              *float64   `json:"annualTask" form:"annualTask" gorm:"column:annual_task;comment:全年任务（万）;size:10;"`                                                            //全年任务（万）
	MonthlyCommissionRating *float64   `json:"monthlyCommissionRating" form:"monthlyCommissionRating" gorm:"column:monthly_commission_rating;comment:月度提成点;size:10;"`                      //月度提成点
	CommissionRatio         *float64   `json:"commissionRatio" form:"commissionRatio" gorm:"column:commission_ratio;comment:提成比例;size:10;"`                                                //提成比例
	Market                  string     `json:"market" form:"market" gorm:"column:market;comment:所责市场;size:100;"`                                                                           //所责市场
	MonthlyShipmentTarget   *float64   `json:"monthlyShipmentTarget" form:"monthlyShipmentTarget" gorm:"column:monthly_shipment_target;comment:月度出货指标;size:10;"`                           //月度出货指标
	ActualShipment          *float64   `json:"actualShipment" form:"actualShipment" gorm:"column:actual_shipment;comment:当月实际出货;size:10;"`                                                 //当月实际出货
	MonthlyCommissionBase   *float64   `json:"monthlyCommissionBase" form:"monthlyCommissionBase" gorm:"column:monthly_commission_base;comment:月度提成基数;size:10;"`                           //月度提成基数
	CompletionRate          *float64   `json:"completionRate" form:"completionRate" gorm:"column:completion_rate;comment:完成率%;size:10;"`                                                   //完成率%
	PerformanceScore        *float64   `json:"performanceScore" form:"performanceScore" gorm:"column:performance_score;comment:绩效考评分数;size:10;"`                                           //绩效考评分数
	AbsenceWork             *float64   `json:"absenceWork" form:"absenceWork" gorm:"column:absence_work;comment:不在岗缺勤;size:10;"`                                                           //不在岗缺勤
	ShippingCommission      *float64   `json:"shippingCommission" form:"shippingCommission" gorm:"column:shipping_commission;comment:出货提成;size:10;"`                                       //出货提成
	PersonalSalesCommission *float64   `json:"personalSalesCommission" form:"personalSalesCommission" gorm:"column:personal_sales_commission;comment:个人销售提成;size:10;"`                     //个人销售提成
	MarketServiceFee        *float64   `json:"marketServiceFee" form:"marketServiceFee" gorm:"column:market_service_fee;comment:市场差补服务费;size:10;"`                                         //市场差补服务费
	MonthlyReward           *float64   `json:"monthlyReward" form:"monthlyReward" gorm:"column:monthly_reward;comment:月度奖励;size:10;"`                                                      //月度奖励
	MonthlyLeaderSubsidy    *float64   `json:"monthlyLeaderSubsidy" form:"monthlyLeaderSubsidy" gorm:"column:monthly_leader_subsidy;comment:品牌负责人月度补贴;size:10;"`                           //品牌负责人月度补贴
	StoreCommission         *float64   `json:"storeCommission" form:"storeCommission" gorm:"column:store_commission;comment:大店开发提成;size:10;"`                                              //大店开发提成
	ProjectNumber           *int       `json:"projectNumber" form:"projectNumber" gorm:"column:project_number;comment:项目数;size:10;"`                                                       //项目数
	HandicraftCosts         *float64   `json:"handicraftCosts" form:"handicraftCosts" gorm:"column:handicraft_costs;comment:手工费;size:10;"`                                                 //手工费
	Performance             *float64   `json:"performance" form:"performance" gorm:"column:performance;comment:业绩;size:10;"`                                                               //业绩
	Commission              *float64   `json:"commission" form:"commission" gorm:"column:commission;comment:提成;size:10;"`                                                                  //提成
	StringValue1            string     `json:"stringValue1" form:"stringValue1" gorm:"column:string_value1;comment:备用字段1;size:200;"`                                                       //备用字段1
	StringValue2            string     `json:"stringValue2" form:"stringValue2" gorm:"column:string_value2;comment:备用字段2;size:200;"`                                                       //备用字段2
	StringValue3            string     `json:"stringValue3" form:"stringValue3" gorm:"column:string_value3;comment:备用字段3;size:200;"`                                                       //备用字段3
	StringValue4            string     `json:"stringValue4" form:"stringValue4" gorm:"column:string_value4;comment:备用字段4;size:200;"`                                                       //备用字段4
	StringValue5            string     `json:"stringValue5" form:"stringValue5" gorm:"column:string_value5;comment:备用字段5;size:200;"`                                                       //备用字段5
	DecimalValue1           *float64   `json:"decimalValue1" form:"decimalValue1" gorm:"column:decimal_value1;comment:备用字段1;size:10;"`                                                     //备用字段1
	DecimalValue2           *float64   `json:"decimalValue2" form:"decimalValue2" gorm:"column:decimal_value2;comment:备用字段2;size:10;"`                                                     //备用字段2
	DecimalValue3           *float64   `json:"decimalValue3" form:"decimalValue3" gorm:"column:decimal_value3;comment:备用字段3;size:10;"`                                                     //备用字段3
	DecimalValue4           *float64   `json:"decimalValue4" form:"decimalValue4" gorm:"column:decimal_value4;comment:备用字段4;size:10;"`                                                     //备用字段4
	DecimalValue5           *float64   `json:"decimalValue5" form:"decimalValue5" gorm:"column:decimal_value5;comment:备用字段5;size:10;"`                                                     //备用字段5
	IntValue1               *int       `json:"intValue1" form:"intValue1" gorm:"column:int_value1;comment:备用字段1;size:10;"`                                                                 //备用字段1
	IntValue2               *int       `json:"intValue2" form:"intValue2" gorm:"column:int_value2;comment:备用字段2;size:10;"`                                                                 //备用字段2
	IntValue3               *int       `json:"intValue3" form:"intValue3" gorm:"column:int_value3;comment:备用字段3;size:10;"`                                                                 //备用字段3
	IntValue4               *int       `json:"intValue4" form:"intValue4" gorm:"column:int_value4;comment:备用字段4;size:10;"`                                                                 //备用字段4
	IntValue5               *int       `json:"intValue5" form:"intValue5" gorm:"column:int_value5;comment:备用字段5;size:10;"`                                                                 //备用字段5
	DateValue1              *time.Time `json:"dateValue1" form:"dateValue1" gorm:"column:date_value1;comment:备用字段1;"`                                                                      //备用字段1
	DateValue2              *time.Time `json:"dateValue2" form:"dateValue2" gorm:"column:date_value2;comment:备用字段2;"`                                                                      //备用字段2
	DateValue3              *time.Time `json:"dateValue3" form:"dateValue3" gorm:"column:date_value3;comment:备用字段3;"`                                                                      //备用字段3
	DateValue4              *time.Time `json:"dateValue4" form:"dateValue4" gorm:"column:date_value4;comment:备用字段4;"`                                                                      //备用字段4
	DateValue5              *time.Time `json:"dateValue5" form:"dateValue5" gorm:"column:date_value5;comment:备用字段5;"`                                                                      //备用字段5
}

// TableName 工资列表 WcStaffSalary自定义表名 wc_staff_salary
func (WcStaffSalary) TableName() string {
	return "wc_staff_salary"
}
