package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//文件上传服务
	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()
	//导入模板
	router.LoadHTMLGlob("template/*")
	//文件上传测试
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})
	router.POST("/upload", uploadController)
	router.Run(":27149")

	/*
		==============================================================DEMO内容
	*/

	////Default returns an Engine instance with the Logger and Recovery middleware already attached.
	//router := gin.Default()
	//// 静态资源加载
	//router.StaticFS("/static", http.Dir("C:\\Users\\97759\\Pictures"))
	//router.StaticFile("/pic2", "C:\\Users\\97759\\Pictures\\2.jpg")
	////导入模板
	//router.LoadHTMLGlob("template/*")
	//router.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200, gin.H{"message": "pong"})
	//})
	////website分组
	//v := router.Group("/")
	//{
	//	v.GET("/index", indexController)
	//	v.GET("about", aboutController)
	//}
	////文件上传测试
	//router.GET("/upload", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "upload.html", nil)
	//})
	//router.POST("/upload", uploadController)
	//
	//router.Run(":27149")

}
func uploadController(context *gin.Context) {
	f, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		context.SaveUploadedFile(f, "./upload/"+f.Filename)
		context.JSON(http.StatusOK, gin.H{"msg": "upload successes"})
	}

}

//func indexController(context *gin.Context) {
//	context.HTML(http.StatusOK, "index.html", gin.H{"Title": "my index"})
//}
//func aboutController(context *gin.Context) {
//	context.HTML(http.StatusOK, "about.html", gin.H{"Title": "my about"})
//}
