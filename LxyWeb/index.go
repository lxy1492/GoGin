package LxyWeb

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index()  {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/html/*")
	// 设置静态文件
	router.StaticFS("/assets/css", http.Dir("./templates/assets/css"))
	router.StaticFS("/assets/fonts", http.Dir("./templates/assets/fonts"))
	router.StaticFS("/assets/js", http.Dir("./templates/assets/js"))
	router.StaticFS("/assets/images/about", http.Dir("./templates/assets/images/about"))
	router.StaticFS("/assets/images/background", http.Dir("./templates/assets/images/background"))
	router.StaticFS("/assets/images/blog", http.Dir("./templates/assets/images/blog"))
	router.StaticFS("/assets/images/brand", http.Dir("./templates/assets/images/brand"))
	router.StaticFS("/assets/images/main-slider", http.Dir("./templates/assets/images/main-slider"))
	router.StaticFS("/assets/images/resources", http.Dir("./templates/assets/images/resources"))
	router.StaticFS("/assets/images/shapes", http.Dir("./templates/assets/images/shapes"))
	router.StaticFS("/assets/images/team", http.Dir("./templates/assets/images/team"))
	router.StaticFS("/assets/images/testimonials", http.Dir("./templates/assets/images/testimonials"))
	// router.StaticFS("/assets/images", http.Dir("./templates/assets/images"))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Lxy WEB")
	})

	// 访问模板
	router.GET("/html/:args", func(c *gin.Context) {
		args := c.Param("args")
		html(c, args)
	})


	router.Run()
}
