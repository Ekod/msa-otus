package handlers

import (
	"github.com/Ekod/msa-otus/app/services/msa-otus-api/handlers/healthgrp"
	"github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Mux(log *zap.SugaredLogger) *gin.Engine {
	router := gin.Default()

	hgh := healthgrp.Handlers{
		Log: log,
	}

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", hgh.ReadinessCheck)
		healthGroup.GET("/liveness", hgh.LivenessCheck)
	}

	return router
}
