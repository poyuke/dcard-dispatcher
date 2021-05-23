package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter append handler to gin engine
func InitRouter(e Env, r *gin.Engine) {
	domain := "*"

	corsInfo := cors.Config{
		AllowOrigins:     []string{domain},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsInfo))

	g := r.Group("/api/v1")
	{
		g.POST("/jobs", e.UploadFileHandler)
		g.GET("/jobs/:id", e.ProcessFileHandler)
		g.GET("/jobs", e.GetJobsListHandler)
	}
}
