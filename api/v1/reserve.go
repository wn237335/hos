package v1

import (
	"github.com/gin-gonic/gin"
	"hospital/model"
	"hospital/utils/errmsg"
	"net/http"
	"strconv"
)

//新增用户预约
func AddReserve(c *gin.Context) {
	var data model.Reserve
	_ = c.ShouldBindJSON(&data)

	code = model.AddReserve(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}


//用户查看预约列表
func GetReserveByUserIdList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	user_id, _ := strconv.Atoi(c.Param("c_user_id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetDoctorByHospitalIdList(user_id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询所有用户预约列表
func GetReserveList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetReserveList(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户预约
func DeleteReserve(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteReserve(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}