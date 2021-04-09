package routes

import (
	"github.com/gin-gonic/gin"
	v1 "hospital/api/v1"
	"hospital/middleware"
	"hospital/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUserList)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		//医院模块
		auth.POST("admin/addHospital", v1.AddHospital)
		auth.PUT("admin/editHospitalById/:id", v1.EditHospitalById)
		auth.DELETE("admin/deleteHospitalById/:id", v1.DeleteHospitalById)

		//部门模块
		auth.POST("admin/addDepartment", v1.AddDepartment)
		auth.PUT("admin/editDepartment/:id", v1.EditDepartment)
		auth.DELETE("admin/deleteDepartment/:id", v1.DeleteDepartment)

		//Doctor模块路由
		auth.POST("admin/addDoctor", v1.AddDoctor)
		auth.PUT("admin/editDoctorInfo/:id", v1.EditDoctorInfo)
		auth.DELETE("admin/deleteDoctor/:id", v1.DeleteDoctor)

		//用户预约模块
		auth.DELETE("admin/deleteReserve/:id", v1.DeleteReserve)

	}


	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
		router.POST("admin/check_token", v1.CheckToken)
		//User模块路由
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)


		//登录模块
		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		//Hospital模块路由
		router.GET("hospital/getHospitalByLevel/:level", v1.GetHospitalByLevel)
		router.GET("hospital/getHospitalInfo/:id", v1.GetHospitalInfo)
		router.GET("hospital/getHospitalList", v1.GetHospitalList)

		//Department模块路由
		router.GET("hospital/getDepartmentInfo/:id", v1.GetDepartmentInfo)
		router.GET("hospital/getDepartmentList", v1.GetDepartmentList)
		router.GET("hospital/getDepartmentListByCID/:cid", v1.GetDepartmentListByCID)

		//Doctor模块路由
		router.GET("hospital/getDoctorInfo/:id", v1.GetDoctorInfo)
		router.GET("hospital/getDoctorList", v1.GetDoctorList)
		router.GET("hospital/getDoctorByHospitalIdList", v1.GetDoctorByHospitalIdList)
		router.GET("hospital/getDoctorByDeparmentIdList", v1.GetDoctorByDeparmentIdList)


		//用户预约模块
		router.POST("user/addReserve", v1.AddReserve)
		router.GET("hospital/getReserveByUserIdList/:id", v1.GetReserveByUserIdList)
		router.GET("hospital/getReserveList", v1.GetReserveList)

	}
	_ = r.Run(utils.HttpPort)
}
