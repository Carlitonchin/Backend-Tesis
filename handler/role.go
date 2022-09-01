package handler

import (
	"log"
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllRoles(ctx *gin.Context) {
	log.Print("entre al handler")
	roles, err := h.RoleService.GetRoles(ctx.Request.Context())
	log.Print("service layer passed")

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}
