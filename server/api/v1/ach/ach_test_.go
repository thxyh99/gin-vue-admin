package ach

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ach"
    achReq "github.com/flipped-aurora/gin-vue-admin/server/model/ach/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AchTestApi struct {
}

var achTestService = service.ServiceGroupApp.AchServiceGroup.AchTestService


// CreateAchTest 创建绩效测试
// @Tags AchTest
// @Summary 创建绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ach.AchTest true "创建绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /achTest/createAchTest [post]
func (achTestApi *AchTestApi) CreateAchTest(c *gin.Context) {
	var achTest ach.AchTest
	err := c.ShouldBindJSON(&achTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := achTestService.CreateAchTest(&achTest); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAchTest 删除绩效测试
// @Tags AchTest
// @Summary 删除绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ach.AchTest true "删除绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achTest/deleteAchTest [delete]
func (achTestApi *AchTestApi) DeleteAchTest(c *gin.Context) {
	ID := c.Query("ID")
	if err := achTestService.DeleteAchTest(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAchTestByIds 批量删除绩效测试
// @Tags AchTest
// @Summary 批量删除绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /achTest/deleteAchTestByIds [delete]
func (achTestApi *AchTestApi) DeleteAchTestByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := achTestService.DeleteAchTestByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAchTest 更新绩效测试
// @Tags AchTest
// @Summary 更新绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ach.AchTest true "更新绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /achTest/updateAchTest [put]
func (achTestApi *AchTestApi) UpdateAchTest(c *gin.Context) {
	var achTest ach.AchTest
	err := c.ShouldBindJSON(&achTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := achTestService.UpdateAchTest(achTest); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAchTest 用id查询绩效测试
// @Tags AchTest
// @Summary 用id查询绩效测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ach.AchTest true "用id查询绩效测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /achTest/findAchTest [get]
func (achTestApi *AchTestApi) FindAchTest(c *gin.Context) {
	ID := c.Query("ID")
	if reachTest, err := achTestService.GetAchTest(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reachTest": reachTest}, c)
	}
}

// GetAchTestList 分页获取绩效测试列表
// @Tags AchTest
// @Summary 分页获取绩效测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query achReq.AchTestSearch true "分页获取绩效测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achTest/getAchTestList [get]
func (achTestApi *AchTestApi) GetAchTestList(c *gin.Context) {
	var pageInfo achReq.AchTestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := achTestService.GetAchTestInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetAchTestPublic 不需要鉴权的绩效测试接口
// @Tags AchTest
// @Summary 不需要鉴权的绩效测试接口
// @accept application/json
// @Produce application/json
// @Param data query achReq.AchTestSearch true "分页获取绩效测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achTest/getAchTestList [get]
func (achTestApi *AchTestApi) GetAchTestPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的绩效测试接口信息",
    }, "获取成功", c)
}
