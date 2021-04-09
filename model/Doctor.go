package model

import (
	"gorm.io/gorm"
	"hospital/utils/errmsg"
)

type Doctor struct {
	gorm.Model
	Id              int    `gorm:"-;primary_key;AUTO_INCREMENT"`
	Name            string `gorm:"type:varchar(20);not null " json:"name" validate:"required,min=4,max=12" label:"姓名"`
	Age             int    `gorm:"type:int" json:"age"  label:"年龄"`
	CHospitalName   string    `gorm:"type:varchar(100)" json:"c_hospital_name"  label:"所属医院名称"`
	CHospitalI      int    `gorm:"type:int" json:"c_hospital_id"  label:"所属医院id"`
	CDepartmentName string   `gorm:"type:varchar(100)" json:"c_department_name"  label:"所在部门名称"`
	CDepartmentId   int    `gorm:"type:int" json:"c_department_id"  label:"所在部门id"`
	Img             string `gorm:"type:varchar(100);not null" json:"img" label:"首图地址"`
	Info            string `gorm:"type:varchar(1000);not null" json:"info" label:"基本信息"`
	ReadCount       int    `gorm:"type:int;not null;default:0" json:"read_count" label:"游览数量"`
}

// 新增医生
func CreateDoc(data *Doctor) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

//  按医院查询医生list
func GetDoctorByHospitalIdList(hospital_id int, pageSize int, pageNum int) ([]Doctor, int, int64) {
	var doc []Doctor
	var total int64
	db.Preload("Doctor").Where("c_hospital_id = ?", hospital_id).Find(&doc).Count(&total)
	err := db.Preload("Doctor").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"c_hospital_id  =?", hospital_id).Find(&doc).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return doc, errmsg.SUCCSE, total
}

//  按部门查询医生list
func GetDoctorByDeparmentIdList(department_id int, pageSize int, pageNum int) ([]Doctor, int, int64) {
	var doc []Doctor
	var total int64
	db.Preload("Doctor").Where("c_department_id = ?", department_id).Find(&doc).Count(&total)
	err := db.Preload("Doctor").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"c_department_id  =?", department_id).Find(&doc).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return doc, errmsg.SUCCSE, total
}

//  查询单个医生
func GetDoctorInfo(id int) (Doctor, int) {
	var doc Doctor
	err := db.Preload("Doctor").Where("id = ?", id).First(&doc).Error
	db.Model(&doc).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return doc, errmsg.ERROR_ART_NOT_EXIST
	}
	return doc, errmsg.SUCCSE
}

//  查询医生列表
func GetDoctorList(pageSize int, pageNum int) ([]Doctor, int, int64) {
	var doc []Doctor
	var err error
	var total int64

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&doc).Error
	// 单独计数
	db.Model(&doc).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return doc, errmsg.SUCCSE, total

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

// 更新医生信息
func EditDoctorInfo(id int, data *Doctor) int {
	var doc Doctor
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
//	maps["age"] = data.Age
//	maps["img"] = data.Img
//	maps["info"] = data.Info

	err = db.Model(&doc).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除医生
func DeleteDoctor(id int) int {
	var doc Doctor
	err = db.Where("id = ? ", id).Delete(&doc).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
