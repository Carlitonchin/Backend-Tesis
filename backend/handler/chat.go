package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getMessagesChat(ctx *gin.Context) {
	question_id := ctx.Param("question_id")

	if question_id == "" {
		type_error := apperrors.BadRequest
		message := "Falto agregar el id de la pregunta"
		err := apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	question_id_uint, err := strconv.ParseUint(question_id, 0, 64)
	if err != nil {
		type_error := apperrors.BadRequest
		message := "el id de la pregunta tiene que ser un entero no negativo"
		err = apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)

		return
	}

	question, err := h.QuestionService.GetQuestionById(ctx.Request.Context(), uint(question_id_uint))

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	if question.UserRefer != user.ID &&
		(question.UserResponsible == nil || *question.UserResponsible != user.ID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("No tiene permisos para acceder a este chat"),
		})

		return
	}

	messages_chat, err := h.ChatService.GetMessages(ctx.Request.Context(), question.ID, user.ID)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"messages": messages_chat,
	})
}
