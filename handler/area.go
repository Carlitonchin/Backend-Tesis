package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/gin-gonic/gin"
)

type areaReq struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) addArea(ctx *gin.Context) {
	var req areaReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	area := &model.Area{
		Name: req.Name,
	}

	area, err := h.AreaService.AddArea(ctx.Request.Context(), area)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"area": area.ID,
	})
}
