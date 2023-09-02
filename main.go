package main

import (
	v1 "GoTODO/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1.InitializeRoutes(r)
	r.Run(":8080")
}
