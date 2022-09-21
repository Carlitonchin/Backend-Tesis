package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/Carlitonchin/Backend-Tesis/some_utils"
	"github.com/gin-gonic/gin"
)

type addQuestionReq struct {
	Text string `json:"text" binding:"required"`
}

func (h *Handler) addQuestion(ctx *gin.Context) {
	var req addQuestionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)

		return
	}

	question := &model.Question{
		Text:      req.Text,
		UserRefer: user.ID,
	}

	question, err = h.QuestionService.AddQuestion(ctx.Request.Context(), question)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"question": question.ID,
	})
}

type clasifyReq struct {
	QuestionId uint `json:"question_id" binding:"required"`
	AreaId     uint `json:"area_id" binding:"required"`
}

func (h *Handler) clasifyQuestion(ctx *gin.Context) {
	var req clasifyReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	err := h.QuestionService.Clasify(ctx.Request.Context(), req.QuestionId, req.AreaId)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

type questionReq struct {
	Question_id uint `json:"question_id" binding:"required"`
}

func (h *Handler) takeQuestion(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.TakeQuestion(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

type responseQuestionReq struct {
	QuestionId uint   `json:"question_id" binding:"required"`
	Response   string `json:"response" binding:"required"`
}

func (h *Handler) responseQuestion(ctx *gin.Context) {
	var req responseQuestionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.ResponseQuestion(ctx, user, req.QuestionId, req.Response)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) upLevel(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.UpLevel(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) upToAdmin(ctx *gin.Context) {
	var req questionReq
	if ok := bindData(ctx, &req); !ok {
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	err = h.QuestionService.UpToAdmin(ctx.Request.Context(), user, req.Question_id)

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getMyQuestions(ctx *gin.Context) {
	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	questions, err := h.QuestionService.GetMyQuestions(ctx.Request.Context(), user.ID)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}

func (h *Handler) getUnclasifiedQuestions(ctx *gin.Context) {
	questions, err := h.QuestionService.GetUnClasifiedQuestions(ctx.Request.Context())

	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}

func (h *Handler) getQuestionByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	if status == "" {
		type_error := apperrors.BadRequest
		message := "Falto agregar el status"
		err := apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	status_uint, err := strconv.ParseUint(status, 0, 64)
	if err != nil {
		type_error := apperrors.BadRequest
		message := "El status tiene que ser un entero no negativo"
		err = apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	status_send, err1 := some_utils.GetUintEnv("STATUS_SEND_CODE")
	status_clasified1, err2 := some_utils.GetUintEnv("STATUS_CLASIFIED1_CODE")
	status_clasified2, err3 := some_utils.GetUintEnv("STATUS_CLASIFIED2_CODE")

	role_student := os.Getenv("ROLE_DEFAULT_STUDENT")
	role_clasifier := os.Getenv("ROLE_DEFAULT_WORKER")
	role_specialist1 := os.Getenv("ROLE_SPECIALIST_LEVEL1")
	role_specialist2 := os.Getenv("ROLE_SPECIALIST_LEVEL2")
	role_admin := os.Getenv("ROLE_ADMIN")

	if some_utils.AtLeastOneError(err1, err2, err3) {
		type_error := apperrors.Internal
		message := "Ocurrio un error inesperado"
		err = apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	user, err := handler_utils.GetUser(ctx)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	role_user := user.Role.Name

	fail := false
	fail = fail || role_user == role_student
	fail = fail || (role_user == role_clasifier && status_uint != uint64(status_send))
	fail = fail || (role_user == role_specialist1 && status_uint != uint64(status_clasified1))
	fail = fail || (role_user == role_specialist2 && status_uint != uint64(status_clasified2))

	if fail {
		type_error := apperrors.Authorization
		message := fmt.Sprintf("No esta autorizado a acceder a las preguntas con status_id = '%v'", status_uint)
		err = apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	var area_id *uint
	if role_admin == role_user {
		area_id = nil
	} else if user.AreaID != nil {
		area_id = user.AreaID
	} else {
		type_error := apperrors.Authorization
		message := fmt.Sprintf("El usuario %s no tiene asignada un ára", user.Name)
		err = apperrors.NewError(type_error, message)
		handler_utils.SendErrorResponse(ctx, err)
		return
	}
	questions, err := h.QuestionService.GetQuestionsByStatus(ctx.Request.Context(), uint(status_uint), area_id)
	if err != nil {
		handler_utils.SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}
