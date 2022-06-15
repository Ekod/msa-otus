package handlers

import (
	"github.com/Ekod/msa-otus/app/services/msa-otus-api/handlers/healthgrp"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Mux(log *zap.SugaredLogger, db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	hgh := healthgrp.Handlers{
		Log:   log,
		DB: db,
	}

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", hgh.ReadinessCheck)
		healthGroup.GET("/liveness", hgh.LivenessCheck)
	}

	return router
}