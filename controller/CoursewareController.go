package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func Upload(c *gin.Context) {
	id,err:=c.Cookie("login")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"message":"请先登录",
		})
	}else {
		if id == "teacher" {
			file,err:=c.FormFile("courseware")
			if err != nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"err:":err.Error(),
				})
				return
			}
			dst:=fmt.Sprintf("./%s",file.Filename)
			c.SaveUploadedFile(file,dst)
			c.JSON(http.StatusOK,gin.H{
				"code":200,
				"message":"上传课件成功",
			})
		}
		if id=="student"{
			c.JSON(http.StatusBadRequest,gin.H{
				"code":400,
				"message":"学生不能上传课件",
			})
		}
	}
}

func Download(c *gin.Context) {
	id,err:=c.Cookie("login")
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"请先登录",
		})
	}else if id=="student"||id=="teacher"{
	file := c.Query("file")
	dst:=fmt.Sprintf("./%s",file)
	//打开文件
	fileTmp, err := os.Open(dst)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"message":"文件未找到",
		})
	}
	defer fileTmp.Close()
	//获取文件的名称
	fileName:=path.Base(dst)
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(dst)
	}
	return
}