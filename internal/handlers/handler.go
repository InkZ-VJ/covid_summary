package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"covid/config"
	"covid/internal/core/ports"
)

type handler struct {
	csvc   ports.CovidService
	router *gin.Engine
}

func NewCovidHdl(csvc ports.CovidService) ports.CovidHandler {
	hdl := &handler{csvc: csvc}
	hdl.setupRouter()
	return hdl
}

func (h *handler) setupRouter() {
	router := gin.Default()
	router.GET("/covid/summary", h.getCovidSummary)

	h.router = router
}

func (h *handler) Start() error {
	return h.router.Run(fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port))
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (h *handler) getCovidSummary(ctx *gin.Context) {
	out, err := h.csvc.GetSummary(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, out)
}
