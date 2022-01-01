package router

import (
	"CSA_Work_Go/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r:=gin.Default()
	StudentRouter:=r.Group("/student")
	{
		StudentRouter.POST("/",controller.StudentPost)
		StudentRouter.GET("/",controller.StudentGet)
		StudentRouter.PUT("/",controller.StudentPut)
		StudentRouter.DELETE("/",controller.StudentDelete)
	}
	TeacherRouter:=r.Group("/teacher")
	{
		TeacherRouter.POST("/",controller.TeacherPost)
		TeacherRouter.GET("/",controller.TeacherGet)
		TeacherRouter.PUT("/",controller.TeacherPut)
		TeacherRouter.DELETE("/",controller.TeacherDelete)
	}
	RecordRouter:=r.Group("/record")
	{
		RecordRouter.POST("/",controller.RecordPost)
		RecordRouter.GET("/",controller.RecordGet)
		RecordRouter.PUT("/",controller.RecordPut)
		RecordRouter.DELETE("/",controller.RecordDelete)
	}
	CoursewareRouter:=r.Group("/courseware")
	{
		CoursewareRouter.GET("/",controller.Download)
		CoursewareRouter.POST("/",controller.Upload)
	}
	return r
}
