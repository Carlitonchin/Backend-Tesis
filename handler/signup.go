package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

type signup_req struct {
	Email    string `json:"email" binding:"required, email"`
	Pass     string `json:"pass" binding:"required,gte=6,lte=30"`
	UserName string `json:"name" binding:"required"`
	Worker   bool   `json:"worker" binding:"required"`
}

func (s *Handler) SignUp(ctx *gin.Context) {
	var req signup_req

	if ok := bindData(ctx, &req); !ok {
		return // error handled at bindData function
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Pass,
		Name:     req.UserName,
		Worker:   req.Worker,
	}

	err := s.UserService.SignUp(ctx, u)

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	token, err := s.TokenService.GetNewPairFromUser(ctx, u, "")

	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"tokens": token,
	})
}
