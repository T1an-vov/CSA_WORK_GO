package controller

import (
	"CSA_Work_Go/module"
	"CSA_Work_Go/service"
	"github.com/gin-gonic/gin"
	"net/http"

)

func RecordPost(c *gin.Context)  {
	var record module.Record
	c.ShouldBind(&record)
	err:=service.RecordPost(&record)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"创建记录成功",
		})
	}
}
func RecordGet(c *gin.Context) {
	var records *[]module.Record
	var err error
	name:=c.Query("name")
	if records,err=service.RecordGet(name);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"records":records,
		})
	}
}
func RecordPut(c *gin.Context) {
	var record module.Record
	c.ShouldBind(&record)
	if err:=service.RecordPut(&record);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"更新记录成功",
		})
	}
}
func RecordDelete(c *gin.Context) {
	var record module.Record
	c.ShouldBind(&record)
	if err:=service.RecordDelete(&record);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"删除记录成功",
		})
	}
}
