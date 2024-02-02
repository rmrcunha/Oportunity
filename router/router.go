package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRouter(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
