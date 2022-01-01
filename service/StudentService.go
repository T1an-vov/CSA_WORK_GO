package service

import (
	"CSA_Work_Go/dao"
	"CSA_Work_Go/module"
	"errors"
)

func StudentGet(name string,password string) error {
	var students []module.Student
	if err:=dao.DB.Where("name=?",name).Find(&students).Error;err!=nil{
		return errors.New("登录")
	}
	for _, student := range students {
		if student.Name==name{
			if student.Password==password{
				return nil
			}
		}
	}
	return errors.New("名称或密码输入错误")
}

func StudentPost(student *module.Student)error{
	if err:=dao.DB.Create(student).Error;err!=nil{
		return errors.New("创建失败")
	}
	return nil
}
func StudentDelete(student *module.Student) error {
	if err:=dao.DB.Where("name=?",student.Name).First(student).Error;err!=nil{
		return errors.New("该学生不存在，删除失败")
	}
	if err:=dao.DB.Delete(student).Error;err!=nil{
		return errors.New("删除失败")
	}
	return nil
}
func StudentPut(name string,password string,newpassword string)error {
	student:=new(module.Student)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(&student).Error;err!=nil{
		return errors.New("名称或密码错误，修改密码失败")
	}
	student.Password=newpassword
	if err:=dao.DB.Save(&student).Error;err!=nil{
		return errors.New("修改密码失败")
	}
	return nil
}