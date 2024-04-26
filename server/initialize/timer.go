package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 同步企微通讯录部门信息 半个小时一次
		//_, err = global.GVA_Timer.AddTaskByFunc("SyncDepartment", "0 */30 * * * ?", func() {
		//	dbEdu := global.GetGlobalDBByDBName("edu")
		//	err := task.SyncDepartment(global.GVA_DB, dbEdu) // 定时任务方法定在task文件包中
		//	if err != nil {
		//		fmt.Println("timer error:", err)
		//	}
		//}, "定时同步容大在线通讯录", option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}

		// 同步企微通讯录员工职位 半个小时一次
		//_, err = global.GVA_Timer.AddTaskByFunc("SyncPosition", "0 */30 * * * ?", func() {
		//	dbEdu := global.GetGlobalDBByDBName("edu")
		//	err := task.SyncPosition(global.GVA_DB, dbEdu) // 定时任务方法定在task文件包中
		//	if err != nil {
		//		fmt.Println("timer error:", err)
		//	}
		//}, "定时同步容大在线通讯录", option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}

		// 同步企微通讯录部门信息 半个小时一次
		//_, err = global.GVA_Timer.AddTaskByFunc("SyncDepartment", "0 */1 * * * ?", func() {
		//	dbEdu := global.GetGlobalDBByDBName("edu")
		//	err := task.SyncStaff(global.GVA_DB, dbEdu) // 定时任务方法定在task文件包中
		//	if err != nil {
		//		fmt.Println("timer error:", err)
		//	}
		//}, "定时同步容大在线通讯录", option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
