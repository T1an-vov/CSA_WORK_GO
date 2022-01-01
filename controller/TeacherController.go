package controller

import (
	"CSA_Work_Go/module"
	"CSA_Work_Go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TeacherGet(c *gin.Context) {
	var teacher module.Teacher
	c.ShouldBind(&teacher)
	if err:=service.TeacherGet(teacher.Name,teacher.Password);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":err.Error(),
		})
	}else {
		cookie := &http.Cookie{
			Name: "login",
			Value: "teacher",
			MaxAge: 3600,
			Path: "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(	http.StatusOK,gin.H{
			"code":200,
			"message":teacher.Name+"登录成功",
		})

	}
}
func TeacherPost(c *gin.Context) {
	var teacher module.Teacher
	c.ShouldBind(&teacher)
	if err:=service.TeacherPost(&teacher);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"创建老师"+teacher.Name+"成功",
		})
	}
}
func TeacherDelete(c *gin.Context) {
	var teacher module.Teacher
	c.ShouldBind(&teacher)
	if err:=service.TeacherDelete(&teacher);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"删除用户"+teacher.Name+"成功",
		})
	}
}
func TeacherPut(c *gin.Context) {
	var teacher module.Teacher
	c.ShouldBind(&teacher)
	newpassword:=c.PostForm("newpassword")
	if err:=service.TeacherPut(teacher.Name,teacher.Password,newpassword);err!=nil{
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


