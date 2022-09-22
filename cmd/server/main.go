package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/config"
)

func main() {
	Config := config.Get()
	config.Init()
	r := gin.Default()
	r.Run(Config.Application.Port)
}
