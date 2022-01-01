package service

import (
	"CSA_Work_Go/dao"
	"CSA_Work_Go/module"
	"errors"
)

func TeacherGet(name string,password string) error {
	var teachers []module.Teacher
	if err:=dao.DB.Where("name=?",name).Find(&teachers).Error;err!=nil{
		return errors.New("查询失败")
	}
	for _, teacher := range teachers {
		if teacher.Name==name{
			if teacher.Password==password{
				return nil
			}
		}
	}
	return errors.New("名称或密码输入错误")
}

func TeacherPost(teacher *module.Teacher)error{
	if err:=dao.DB.Create(teacher).Error;err!=nil{
		return errors.New("创建失败")
	}
	return nil
}
func TeacherDelete(teacher *module.Teacher) error {
	if err:=dao.DB.Where("name=?",teacher.Name).First(teacher).Error;err!=nil{
		return errors.New("该老师不存在，删除失败")
	}
	if err:=dao.DB.Delete(teacher).Error;err!=nil{
		return errors.New("删除失败")
	}
	return nil
}
func TeacherPut(name string,password string,newpassword string)error {
	teacher:=new(module.Teacher)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(&teacher).Error;err!=nil{
		return errors.New("名称或密码错误，修改密码失败")
	}
	teacher.Password=newpassword
	if err:=dao.DB.Save(&teacher).Error;err!=nil{
		return errors.New("修改密码失败")
	}
	return nil
}