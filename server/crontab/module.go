package crontab

import (
	"github.com/robfig/cron/v3"
)

type Module struct {
}

func (m Module) Run() {
	s := NewService()
	c := cron.New()
	//{secondParser, "0 5 * * * *", every5min(time.Local)},
	//{standardParser, "5 * * * *", every5min(time.Local)},
	// 同步企微部门列表(每10分钟执行一次)(接口请求受限改成从edu同步)
	_, err := c.AddFunc("*/10 * * * ?", s.SyncDepartment)
	if err != nil {
		return
	}

	//启动定时任务
	c.Start()
}
