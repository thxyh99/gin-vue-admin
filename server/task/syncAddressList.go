package task

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Department struct {
	DepartmentID int
	Parentid     int
	Name         string
	Status       int
}

type NewDepartment struct {
	ID           int
	DepartmentID int
	Parentid     int
	Name         string
	NameEn       string
	Order        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func SyncDepartment(db, dbEdu *gorm.DB) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}

	if dbEdu == nil {
		return errors.New("db-edu Cannot be empty")
	}

	//定时任务刷一个小时内有更新的数据
	timestamp := time.Now().Unix() - 3600
	//初始化全部数据
	//timestamp := 0

	// 执行原生SQL查询
	var departments []Department
	sqlEdu := "SELECT did AS department_id, parentid, `name`,`status` FROM g_department_pro WHERE update_at > ? AND did > ? ORDER BY did ASC"
	dbEdu.Raw(sqlEdu, timestamp, 0).Scan(&departments)

	fmt.Println("SQL query:", time.Now(), timestamp)

	var item NewDepartment
	// 将查询结果插入到新表中
	for _, department := range departments {

		fmt.Println(department)
		now := time.Now()
		newDepartment := NewDepartment{
			DepartmentID: department.DepartmentID,
			Parentid:     department.Parentid,
			Name:         department.Name,
			NameEn:       "",
			Order:        0,
			UpdatedAt:    now,
		}
		if department.Status == 0 {
			newDepartment.DeletedAt = &now
		} else {
			newDepartment.DeletedAt = nil
		}

		itemResult := db.Table("wc_department").Where("department_id=?", department.DepartmentID).Limit(1).Find(&item)
		if itemResult.Error != nil {
			if errors.Is(itemResult.Error, gorm.ErrRecordNotFound) {
				// 如果记录不存在，直接跳过
			} else {
				// 其他错误情况
				return itemResult.Error
			}
		}

		if item.ID > 0 {
			result := db.Table("wc_department").Where("id=?", item.ID).Updates(&newDepartment)
			if result.Error != nil {
				fmt.Println("Failed to update department:", result.Error)
			}
		} else {
			newDepartment.CreatedAt = now
			result := db.Table("wc_department").Create(&newDepartment)
			if result.Error != nil {
				fmt.Println("Failed to insert department:", result.Error)
			}
		}

	}

	return nil
}
