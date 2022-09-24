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
	addDentistRoutes(v1)
	addAppointmentRoutes(v1)
	return r
}

func addPatientRoutes(rg *gin.RouterGroup) {
	patient := rg.Group("/patient")
	han := handler.NewPatientHandler()

	patient.GET("/:id", han.PGetByID())
	patient.POST("/", han.PAdd())
	patient.PUT("/", han.PUpdate())
	patient.DELETE("/:id", han.PDelete())
	patient.PATCH("/address", han.PChangeAddresById())
}

func addDentistRoutes(rg *gin.RouterGroup) {
	dentist := rg.Group("/dentist")
	han := handler.NewDentistHandler()

	dentist.GET("/:id", han.DGetByID())
	dentist.POST("/", han.DAdd())
	dentist.PUT("/", han.DUpdate())
	dentist.DELETE("/:id", han.DDelete())
	dentist.PATCH("/address", han.DChangeRegistrationIdById())
}

func addAppointmentRoutes(rg *gin.RouterGroup) {
	appointment := rg.Group("/appointment")
	han := handler.NewAppointmentHandler()

	appointment.GET("/:id", han.AGetByID())
	appointment.POST("/", han.AAdd())
	appointment.PUT("/", han.AUpdate())
	appointment.DELETE("/:id", han.ADelete())
	appointment.PATCH("/time", han.AChangeTimeById())
}
