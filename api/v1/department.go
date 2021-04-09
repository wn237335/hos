package v1

import (
	"github.com/gin-gonic/gin"
	"hospital/model"
	"hospital/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加部门
func AddDepartment(c *gin.Context) {
	var data model.Department
	_ = c.ShouldBindJSON(&data)
	code = model.CheckDepartment(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateDepartment(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询部门信息
func GetDepartmentInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetDepartmentInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// 查询部门列表
func GetDepartmentList(c *gin.Context) {
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

	data, total := model.GetDepartmentList(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//查询部门列表by医院id
func GetDepartmentListByCID(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetDepartmentListByCID(cid, pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询单个
//func GetCateInfo(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data,code := model.GetCateInfo(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// 编辑部门名子
func EditDepartment(c *gin.Context) {
	var data model.Department
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckDepartment(data.Name)
	if code == errmsg.SUCCSE {
		model.EditDepartment(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 删除部门
func DeleteDepartment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteDepartment(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
