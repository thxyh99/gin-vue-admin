package task

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
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

type Position struct {
	Name string
	DEL  int
}

type NewPosition struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Staff struct {
	Uid          int
	Userid       string
	DepartmentId int
	Realname     string
	Mobile       string
	IdNumber     string
	SsNumber     string
	Position     string
	Status       int
	Identity     int
	Del          int
}

type NewStaff struct {
	ID               int
	Userid           string
	Name             string
	Department       string
	Order            int
	PositionId       int
	Position         string
	Gender           int
	Email            string
	BizMail          string
	IsLeaderInDept   string
	DirectLeader     string
	Avatar           string
	ThumbAvatar      string
	Telephone        string
	Alias            string
	Address          string
	OpenUserid       string
	MainDepartment   int
	Extattr          string
	Status           int
	QrCode           string
	ExternalPosition string
	ExternalProfile  string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
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

func SyncPosition(db, dbEdu *gorm.DB) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}

	if dbEdu == nil {
		return errors.New("db-edu Cannot be empty")
	}

	//定时任务刷一个小时内有更新的数据
	//timestamp := time.Now().Unix() - 3600
	//初始化全部数据
	timestamp := 0

	// 执行原生SQL查询
	var positions []Position
	sqlEdu := "SELECT `name`,`del` FROM g_position_pro WHERE update_at > ? AND `type` = ?"
	dbEdu.Raw(sqlEdu, timestamp, 1).Scan(&positions)

	fmt.Println("SQL query:", time.Now(), timestamp)

	var item NewPosition
	// 将查询结果插入到新表中
	for _, position := range positions {

		fmt.Println(position)
		now := time.Now()
		newPosition := NewPosition{
			Name:      position.Name,
			UpdatedAt: now,
		}
		if position.DEL == 1 {
			newPosition.DeletedAt = &now
		} else {
			newPosition.DeletedAt = nil
		}

		itemResult := db.Table("wc_position").Where("name=?", position.Name).Limit(1).Find(&item)
		if itemResult.Error != nil {
			if errors.Is(itemResult.Error, gorm.ErrRecordNotFound) {
				// 如果记录不存在，直接跳过
			} else {
				// 其他错误情况
				return itemResult.Error
			}
		}

		if item.ID > 0 {
			result := db.Table("wc_position").Where("id=?", item.ID).Updates(&newPosition)
			if result.Error != nil {
				fmt.Println("Failed to update position:", result.Error)
			}
		} else {
			newPosition.CreatedAt = now
			result := db.Table("wc_position").Create(&newPosition)
			if result.Error != nil {
				fmt.Println("Failed to insert position:", result.Error)
			}
		}

	}

	return nil
}

func SyncStaff(db, dbEdu *gorm.DB) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}

	if dbEdu == nil {
		return errors.New("db-edu Cannot be empty")
	}

	//定时任务刷一个小时内有更新的数据
	//timestamp := time.Now().Unix() - 3600
	//初始化全部数据
	timestamp := 0

	// 执行原生SQL查询
	var staffs []Staff
	sqlEdu := "SELECT * FROM g_staff_pro WHERE update_at > ? AND identity < ?"
	dbEdu.Raw(sqlEdu, timestamp, 3).Scan(&staffs)

	fmt.Println("SQL query:", time.Now(), timestamp)

	var item NewStaff
	// 将查询结果插入到新表中
	for _, staff := range staffs {

		fmt.Println(staff)
		now := time.Now()
		newStaff := NewStaff{
			Userid:           staff.Userid,
			Name:             staff.Realname,
			Department:       strconv.Itoa(staff.DepartmentId),
			Order:            0,
			PositionId:       0,
			Position:         staff.Position,
			Gender:           0,
			Email:            "",
			BizMail:          "",
			IsLeaderInDept:   "",
			DirectLeader:     "",
			Avatar:           "",
			ThumbAvatar:      "",
			Telephone:        "",
			Alias:            "",
			Address:          "",
			OpenUserid:       "",
			MainDepartment:   0,
			Extattr:          "",
			Status:           staff.Status,
			QrCode:           "",
			ExternalPosition: "",
			ExternalProfile:  "",
			UpdatedAt:        now,
		}
		if staff.Status == 0 {
			newStaff.DeletedAt = &now
		} else {
			newStaff.DeletedAt = nil
		}

		itemResult := db.Table("wc_staff").Where("userid=?", staff.Userid).Limit(1).Find(&item)
		if itemResult.Error != nil {
			if errors.Is(itemResult.Error, gorm.ErrRecordNotFound) {
				// 如果记录不存在，直接跳过
			} else {
				// 其他错误情况
				return itemResult.Error
			}
		}

		if item.ID > 0 {
			result := db.Table("wc_staff").Where("id=?", item.ID).Updates(&newStaff)
			if result.Error != nil {
				fmt.Println("Failed to update staff:", result.Error)
			}
		} else {
			newStaff.CreatedAt = now
			result := db.Table("wc_staff").Create(&newStaff)
			if result.Error != nil {
				fmt.Println("Failed to insert staff:", result.Error)
			}
		}

	}

	return nil
}
