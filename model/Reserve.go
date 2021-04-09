package model

import (
	"gorm.io/gorm"
	"hospital/utils/errmsg"
)

type Reserve struct {
	gorm.Model
	Id        int    `gorm:"-;primary_key;AUTO_INCREMENT"`
	CUserId int `gorm:"type:int;not null " json:"c_user_id" label:"用户id"`
	CUserName string `gorm:"type:int;not null " json:"c_user_name" label:"用户名"`
	CHospitalName int  `gorm:"type:int" json:"c_hospital_name"  label:"所属医院"`
	CHospitalId int  `gorm:"type:int" json:"c_hospital_id"  label:"所属id"`
	CDepartmentName  int `gorm:"type:int" json:"c_department_name"  label:"所在部门"`
	CDepartmentId  int `gorm:"type:int" json:"c_department_id"  label:"所在部门id"`
	CDoctorName  int `gorm:"type:int" json:"c_doctor_name"  label:"挂号医生姓名"`
	CDoctorId  int `gorm:"type:int" json:"c_doctor_id"  label:"挂号医生id"`
}


//用户预约
func AddReserve (data *Reserve) int{
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

//用户查看预约列表
func GetUserReserveList(user_id int, pageSize int, pageNum int) ([]Reserve, int, int64) {
	var res []Reserve
	var total int64
	db.Preload("Reserve").Where("c_user_id = ?", user_id).Find(&res).Count(&total)
	err := db.Preload("Doctor").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"c_user_id  =?", user_id).Find(&res).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return res, errmsg.SUCCSE, total
}

//查询所有用户预约列表
func GetReserveList(pageSize int, pageNum int) ([]Reserve, int, int64) {
	var res []Reserve
	var err error
	var total int64

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&res).Error
	// 单独计数
	db.Model(&res).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return res, errmsg.SUCCSE, total

	/*if title == "" {

	}
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?",
		title+"%",
	).Find(&HospitalList).Error
	// 单独计数
	db.Model(&HospitalList).Where("title LIKE ?", title+"%").Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return HospitalList, errmsg.SUCCSE, total*/
}


//编辑用户预约


//删除用户预约
func DeleteReserve(id int) int {
	var res Reserve
	err = db.Where("id = ? ", id).Delete(&res).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}





