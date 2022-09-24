package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/config"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/db"
)

func main() {
	Config := config.Get()
	db.Test()
	r := gin.Default()
	router(r, 1)
	r.Run(Config.Application.Port)
}
