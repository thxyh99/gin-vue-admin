// 自动生成模板WcStaffSalary
package weChat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 薪资奖金 结构体  WcStaffSalary
type WcStaffSalary struct {
 global.GVA_MODEL 
      StaffId  *int `json:"staffId" form:"staffId" gorm:"column:staff_id;comment:员工ID;size:10;" binding:"required"`  //员工ID 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:200;" binding:"required"`  //姓名 
      EmploymentDate  *time.Time `json:"employmentDate" form:"employmentDate" gorm:"column:employment_date;comment:入职时间;" binding:"required"`  //入职时间 
      Deadline  *time.Time `json:"deadline" form:"deadline" gorm:"column:deadline;comment:截止日期;" binding:"required"`  //截止日期 
      CompanyAge  *int `json:"companyAge" form:"companyAge" gorm:"column:company_age;comment:司龄;size:3;" binding:"required"`  //司龄 
      SalaryDays  *int `json:"salaryDays" form:"salaryDays" gorm:"column:salary_days;comment:发薪天数;size:3;" binding:"required"`  //发薪天数 
      Overtime  *float64 `json:"overtime" form:"overtime" gorm:"column:overtime;comment:加班情况;size:5;" binding:"required"`  //加班情况 
      Absence  *float64 `json:"absence" form:"absence" gorm:"column:absence;comment:缺勤情况;size:5;" binding:"required"`  //缺勤情况 
      Expense  string `json:"expense" form:"expense" gorm:"column:expense;comment:费用科目;size:100;" binding:"required"`  //费用科目 
      Position  string `json:"position" form:"position" gorm:"column:position;comment:职位;size:200;" binding:"required"`  //职位 
      Education  string `json:"education" form:"education" gorm:"column:education;comment:学历;size:100;" binding:"required"`  //学历 
      GradeSalary  *float64 `json:"gradeSalary" form:"gradeSalary" gorm:"column:grade_salary;comment:等级工资;size:10;" binding:"required"`  //等级工资 
      SkillSalary  *float64 `json:"skillSalary" form:"skillSalary" gorm:"column:skill_salary;comment:职称技能津贴;size:10;" binding:"required"`  //职称技能津贴 
      AllowancesSalary  *float64 `json:"allowancesSalary" form:"allowancesSalary" gorm:"column:allowances_salary;comment:津贴;size:10;" binding:"required"`  //津贴 
      AgeAllowanceSalary  *float64 `json:"ageAllowanceSalary" form:"ageAllowanceSalary" gorm:"column:age_allowance_salary;comment:司龄津贴;size:10;" binding:"required"`  //司龄津贴 
      EducationSalary  *float64 `json:"educationSalary" form:"educationSalary" gorm:"column:education_salary;comment:学历工资;size:10;" binding:"required"`  //学历工资 
      FullSalary  *float64 `json:"fullSalary" form:"fullSalary" gorm:"column:full_salary;comment:全勤;size:10;" binding:"required"`  //全勤 
      FilialSalary  *float64 `json:"filialSalary" form:"filialSalary" gorm:"column:filial_salary;comment:尽孝基金;size:10;" binding:"required"`  //尽孝基金 
      PerformanceSalary  *float64 `json:"performanceSalary" form:"performanceSalary" gorm:"column:performance_salary;comment:业绩奖;size:10;" binding:"required"`  //业绩奖 
      FestivalSalary  *float64 `json:"festivalSalary" form:"festivalSalary" gorm:"column:festival_salary;comment:节日金;size:10;" binding:"required"`  //节日金 
      BirthdaySalary  *float64 `json:"birthdaySalary" form:"birthdaySalary" gorm:"column:birthday_salary;comment:生日金;size:10;" binding:"required"`  //生日金 
      JobSubsidy  *float64 `json:"jobSubsidy" form:"jobSubsidy" gorm:"column:job_subsidy;comment:岗位补助;size:10;" binding:"required"`  //岗位补助 
      LeaderSubsidy  *float64 `json:"leaderSubsidy" form:"leaderSubsidy" gorm:"column:leader_subsidy;comment:组长补助;size:10;" binding:"required"`  //组长补助 
      TelephoneSubsidy  *float64 `json:"telephoneSubsidy" form:"telephoneSubsidy" gorm:"column:telephone_subsidy;comment:电话补助;size:10;" binding:"required"`  //电话补助 
      DormitorySubsidy  *float64 `json:"dormitorySubsidy" form:"dormitorySubsidy" gorm:"column:dormitory_subsidy;comment:宿舍补助;size:10;" binding:"required"`  //宿舍补助 
      TransportSubsidy  *float64 `json:"transportSubsidy" form:"transportSubsidy" gorm:"column:transport_subsidy;comment:交通补助;size:10;" binding:"required"`  //交通补助 
      MealSubsidy  *float64 `json:"mealSubsidy" form:"mealSubsidy" gorm:"column:meal_subsidy;comment:餐补;size:10;" binding:"required"`  //餐补 
      TravelSubsidy  *float64 `json:"travelSubsidy" form:"travelSubsidy" gorm:"column:travel_subsidy;comment:出差补助;size:10;" binding:"required"`  //出差补助 
      PackagingSubsidy  *float64 `json:"packagingSubsidy" form:"packagingSubsidy" gorm:"column:packaging_subsidy;comment:分包装间;size:10;" binding:"required"`  //分包装间 
      DormSubsidy  *float64 `json:"dormSubsidy" form:"dormSubsidy" gorm:"column:dorm_subsidy;comment:宿舍费;size:10;" binding:"required"`  //宿舍费 
      OtherSubsidy  *float64 `json:"otherSubsidy" form:"otherSubsidy" gorm:"column:other_subsidy;comment:其他(补助);size:10;" binding:"required"`  //其他(补助) 
      TotalSubsidy  *float64 `json:"totalSubsidy" form:"totalSubsidy" gorm:"column:total_subsidy;comment:合计(补助);size:10;" binding:"required"`  //合计(补助) 
      OvertimeFee  *float64 `json:"overtimeFee" form:"overtimeFee" gorm:"column:overtime_fee;comment:加班费;size:10;" binding:"required"`  //加班费 
      AbsenceFee  *float64 `json:"absenceFee" form:"absenceFee" gorm:"column:absence_fee;comment:缺勤;size:10;" binding:"required"`  //缺勤 
      TotalFee  *float64 `json:"totalFee" form:"totalFee" gorm:"column:total_fee;comment:合计(加班及缺勤);size:10;" binding:"required"`  //合计(加班及缺勤) 
      ShouldPay  *float64 `json:"shouldPay" form:"shouldPay" gorm:"column:should_pay;comment:应发工资;size:10;" binding:"required"`  //应发工资 
      PensionPay  *float64 `json:"pensionPay" form:"pensionPay" gorm:"column:pension_pay;comment:养老保险;size:10;" binding:"required"`  //养老保险 
      MedicalPay  *float64 `json:"medicalPay" form:"medicalPay" gorm:"column:medical_pay;comment:医疗保险;size:10;" binding:"required"`  //医疗保险 
      UnemploymentPay  *float64 `json:"unemploymentPay" form:"unemploymentPay" gorm:"column:unemployment_pay;comment:失业保险;size:10;" binding:"required"`  //失业保险 
      HousingPay  *float64 `json:"housingPay" form:"housingPay" gorm:"column:housing_pay;comment:住房公积金;size:10;" binding:"required"`  //住房公积金 
      PreFilialSalary  *float64 `json:"preFilialSalary" form:"preFilialSalary" gorm:"column:pre_filial_salary;comment:预存尽孝基金;size:10;" binding:"required"`  //预存尽孝基金 
      Other  *float64 `json:"other" form:"other" gorm:"column:other;comment:其他;size:10;" binding:"required"`  //其他 
      Sent  *float64 `json:"sent" form:"sent" gorm:"column:sent;comment:已发;size:10;" binding:"required"`  //已发 
      ActualPay  *float64 `json:"actualPay" form:"actualPay" gorm:"column:actual_pay;comment:实发小计;size:10;" binding:"required"`  //实发小计 
}


// TableName 薪资奖金 WcStaffSalary自定义表名 wc_staff_salary
func (WcStaffSalary) TableName() string {
  return "wc_staff_salary"
}

