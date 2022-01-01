package service

import (
"CSA_Work_Go/dao"
"CSA_Work_Go/module"
"errors"
)

func RecordGet(name string) (*[]module.Record,error) {
	var records []module.Record
	if err:=dao.DB.Where("name=?",name).Find(&records).Error;err!=nil{
		return nil,errors.New("该记录不存在，查询失败")
	}
	return &records,nil
}

func RecordPost(record *module.Record)error{
	if err:=dao.DB.Create(record).Error;err!=nil{
		return errors.New("该记录已存在，创建失败")
	}
	return nil
}
func RecordDelete(record *module.Record) error {
	if err:=dao.DB.Where("name=? and course=?",record.Name,record.Course).First(record).Error;err!=nil{
		return errors.New("该记录不存在，删除失败")
	}
	if err:=dao.DB.Delete(record).Error;err!=nil{
		return errors.New("删除失败")
	}
	return nil
}
func RecordPut( newrecord *module.Record) error {
	var record module.Record
	if err:=dao.DB.Where("name=? and course=?",record.Name,record.Course).First(&record).Error;err!=nil{
		return errors.New("记录不存在，更新失败")
	}
	if err:=dao.DB.Model(&record).Update(newrecord).Error;err!=nil{
		return err
	}
	return nil
}