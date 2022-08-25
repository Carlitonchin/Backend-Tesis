package handler

import (
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InvalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func bindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		var invalidArgs []InvalidArgument

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {
				invalidArgs = append(invalidArgs, InvalidArgument{
					Field: err.Field(),
					Value: err.Value().(string),
					Tag:   err.Tag(),
					Param: err.Param(),
				})
			}

			message := "Request invalido, mirar 'invalidArgs'"
			type_error := apperrors.BadRequest

			err := apperrors.NewError(type_error, message)

			ctx.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})

			return false
		}
	}

	return true
}
