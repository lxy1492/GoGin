package LxyWeb

import (
	"github.com/gin-gonic/gin"
)

func html(c *gin.Context, name string)  {
	file_path := name
	// fmt.Println(file_path)
	c.HTML(200, file_path, "flysnow_org")
}
