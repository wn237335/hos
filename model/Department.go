package model

import (
	"gorm.io/gorm"
	"hospital/utils/errmsg"
)

type Department struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	CHospital int `gorm:"type:int" json:"c_hospital"`
	CName   string `gorm:"type:varchar(200)" json:"c_name"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询部门是否存在
func CheckDepartment(name string) (code int) {
	var dep Department
	db.Select("id").Where("name = ?", name).First(&dep)
	if dep.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// 新增部门
func CreateDepartment(data *Department) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询单个部门信息
func GetDepartmentInfo(id int) (Department, int) {
	var dep Department
	db.Where("id = ?", id).First(&dep)
	return dep, errmsg.SUCCSE
}

// 查询部门列表
func GetDepartmentList(pageSize int, pageNum int) ([]Department, int64) {
	var dep []Department
	var total int64
	err = db.Find(&dep).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&dep).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return dep, total
}

// 查询部门列表
func GetDepartmentListByCID(cid,pageSize int, pageNum int) ([]Department, int64) {
	var dep []Department
	var total int64
	err = db.Model(&Department{}).Where("c_hospital = ?", cid).Find(&dep).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&dep).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return dep, total
}


// 编辑部门信息
func EditDepartment(id int, data *Department) int {
	var dep Department
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&dep).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除部门
func DeleteDepartment(id int) int {
	var dep Department
	err = db.Where("id = ? ", id).Delete(&dep).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
