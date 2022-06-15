package healthgrp

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Ekod/msa-otus/sys/database"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Handlers struct {
	Log *zap.SugaredLogger
	DB  *sqlx.DB
}

func (h *Handlers) LivenessCheck(c *gin.Context) {
	status := "OK"
	statusCode := http.StatusOK
	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	if err := responseCreator(c.Writer, statusCode, data); err != nil {
		h.Log.Errorw("liveness", "ERROR", err)
	}

	h.Log.Infow("liveness", "statusCode", statusCode, "method", c.Request.Method, "path", c.Request.URL.Path, "remoteaddr", c.Request.RemoteAddr)
}

func (h *Handlers) ReadinessCheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	status := "OK"
	statusCode := http.StatusOK

	if err := database.StatusCheck(ctx, h.DB); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}
	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	if err := responseCreator(c.Writer, statusCode, data); err != nil {
		h.Log.Errorw("readiness", "ERROR", err)
	}

	h.Log.Infow("readiness", "statusCode", statusCode, "method", c.Request.Method, "path", c.Request.URL.Path, "remoteaddr", c.Request.RemoteAddr)
}

func responseCreator(w http.ResponseWriter, statusCode int, data any) error {

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
