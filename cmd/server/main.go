package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/cmd/server/handler"
	"github.com/sagarciaescobar/integrador-final-backend-III/config"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
)

func main() {
	Config := config.Get()
	db.Test()
	r := gin.Default()
	r.GET("/api/patient/:id", handler.NewPatientHandler().GetByID())
	r.Run(Config.Application.Port)
}
