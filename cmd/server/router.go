package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/cmd/server/handler"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/web"
)

func router(r *gin.Engine, version int) *gin.Engine {
	api := r.Group("/api")
	v1 := api.Group("/v" + strconv.Itoa(version))
	v1.GET("/test", func(c *gin.Context) {
		web.Success(c, 200, "ping")
	})
	addPatientRoutes(v1)
	addPatientRoutes(v1)
	return r
}

func addPatientRoutes(rg *gin.RouterGroup) {
	patient := rg.Group("/patient")
	han := handler.NewPatientHandler()

	patient.GET("/:id", han.GetByID())
	patient.POST("/", han.Add())
	patient.PUT("/", han.Update())
	patient.DELETE("/:id", han.Delete())
	patient.PATCH("/address", han.ChangeAddresById())
}
