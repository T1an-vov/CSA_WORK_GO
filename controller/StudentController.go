package controller

import (
	"CSA_Work_Go/module"
	"CSA_Work_Go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentGet(c *gin.Context) {
	var student module.Student
	c.ShouldBind(&student)
	if err:=service.StudentGet(student.Name,student.Password);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":err.Error(),
		})
	}else {
		cookie := &http.Cookie{
			Name: "login",
			Value: "student",
			MaxAge: 3600,
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(	http.StatusOK,gin.H{
			"code":200,
			"message":student.Name+"登录成功",
		})

	}
}
func StudentPost(c *gin.Context) {
	var student module.Student
	c.ShouldBind(&student)
	if err:=service.StudentPost(&student);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"创建学生"+student.Name+"成功",
		})
	}
}
func StudentDelete(c *gin.Context) {
	var student module.Student
	c.ShouldBind(&student)
	if err:=service.StudentDelete(&student);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"删除用户"+student.Name+"成功",
		})
	}
}
func StudentPut(c *gin.Context) {
	var student module.Student
	c.ShouldBind(&student)
	newpassword:=c.PostForm("newpassword")
	if err:=service.StudentPut(student.Name,student.Password,newpassword);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"修改密码成功",
		})
	}
}


