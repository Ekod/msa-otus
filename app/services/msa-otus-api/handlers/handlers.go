package handlers

import (
	"github.com/Ekod/msa-otus/app/services/msa-otus-api/handlers/healthgrp"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Mux(log *zap.SugaredLogger, db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	hgh := healthgrp.Handlers{
		Log: log,
		DB:  db,
	}

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", hgh.ReadinessCheck)
		healthGroup.GET("/liveness", hgh.LivenessCheck)
	}

	return router
}
