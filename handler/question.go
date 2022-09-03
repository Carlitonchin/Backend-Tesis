package handler

import (
	"net/http"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/some_utils"
	"github.com/gin-gonic/gin"
)

type questionReq struct {
	Text string `json:"text" binding:"required"`
}

func (h *Handler) addQuestion(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)

		return
	}

	status_id, err := some_utils.GetUintEnv("STATUS_SEND_CODE")

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)

		return
	}

	question := &model.Question{
		Text:      req.Text,
		UserRefer: user.ID,
		StatusId:  status_id,
	}

	question, err = h.UserService.AddQuestion(ctx.Request.Context(), question)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"question": question.ID,
	})
}
