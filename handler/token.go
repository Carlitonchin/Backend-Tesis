package handler

import (
	"log"
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens handler
func (h *Handler) tokens(c *gin.Context) {
	// bind JSON to req of type tokensRew
	var req tokensReq

	if ok := bindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	// verify refresh JWT
	refreshToken, err := h.TokenService.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// get up-to-date user
	u, err := h.UserService.GetById(ctx, refreshToken.UID)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.GetNewPairFromUser(ctx, u, refreshToken.ID)

	if err != nil {
		log.Printf("Failed to create tokens for user: %+v. Error: %v\n", u, err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}

type updateRoleReq struct {
	UserId uint `json:"user_id" binding:"required"`
	RoleId uint `json:"role_id" binding:"required"`
}

func (h *Handler) updateUserRole(ctx *gin.Context) {
	var req updateRoleReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	err := h.UserService.AddRoleToUser(ctx.Request.Context(), req.UserId, req.RoleId)

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
