package model

import (
	"gorm.io/gorm"
	"hospital/utils/errmsg"
)

type Hospital struct {
	gorm.Model
	ID                uint   `gorm:"primary_key;auto_increment" json:"id"`
	EstablishmentTime string `gorm:"type:varchar(100);" json:"establishment_time" label:"医院建立时间"`
	Name              string `gorm:"type:varchar(100);" json:"name" label:"医院名称"`
	Telephone         string `gorm:"type:varchar(100);not null" json:"telephone" label:"电话号码"`
	Level             string `gorm:"type:varchar(100);not null" json:"level" label:"医院级别如三甲医院"`
	Email             string `gorm:"type:varchar(100);not null" json:"email" label:"邮箱"`
	Address           string `gorm:"type:varchar(100);not null" json:"address" label:"地址"`
	Introduction      string `gorm:"type:varchar(100);not null" json:"introduction" label:"医院简介"`
	Img               string `gorm:"type:varchar(100);not null" json:"img" label:"首图地址"`
	Info              string `gorm:"type:varchar(1000);not null" json:"info" label:"基本信息"`
	ReadCount         int    `gorm:"type:int;not null;default:0" json:"read_count" label:"游览数量"`
}

// 新增医院
func CreateHospital(data *Hospital) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

//  按医院级别查询医院
func GetHospitalByLevel(level string, pageSize int, pageNum int) ([]Hospital, int, int64) {
	var hospitalList []Hospital
	var total int64
	db.Preload("Hospital").Where("level = ?", level).Find(&hospitalList).Count(&total)
	err := db.Preload("Hospital").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"level =?", level).Find(&hospitalList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return hospitalList, errmsg.SUCCSE, total
}

//  查询单个医院
func GetHospitalInfo(id int) (Hospital, int) {
	var hos Hospital
	err := db.Model(&Hospital{}).Where("id = ?", id).Last(&hos).Error
	//err := db.Preload("Hospital").Where("id = ?", id).First(&hos).Error
	db.Model(&hos).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return hos, errmsg.ERROR_ART_NOT_EXIST
	}
	return hos, errmsg.SUCCSE
}

//  查询医院列表
func GetHospitalList(pageSize int, pageNum int) ([]Hospital, int, int64) {
	var HospitalList []Hospital
	var err error
	var total int64

	//	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&HospitalList).Error
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&HospitalList).Error
	// 单独计数
	db.Model(&HospitalList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return HospitalList, errmsg.SUCCSE, total
	/*if title == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&HospitalList).Error
		// 单独计数
		db.Model(&HospitalList).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return HospitalList, errmsg.SUCCSE, total
	}*/
	/*err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?",
		title+"%",
	).Find(&HospitalList).Error
	// 单独计数
	db.Model(&HospitalList).Where("title LIKE ?", title+"%").Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return HospitalList, errmsg.SUCCSE, total*/
}

//更新医院信息
func EditHospitalById(id int, data *Hospital) int {
	var hos Hospital
	var maps = make(map[string]interface{})
	maps["establishment_time"] = data.EstablishmentTime
	maps["name"] = data.Name
	maps["telephone"] = data.Telephone
	maps["level"] = data.Level
	maps["email"] = data.Email
	maps["address"] = data.Address
	maps["introduction"] = data.Introduction
	maps["img"] = data.Img
	maps["info"] = data.Info

	err = db.Model(&hos).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除医院
func DeleteHospitalById(id int) int {
	var hos Hospital
	err = db.Where("id = ? ", id).Delete(&hos).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
