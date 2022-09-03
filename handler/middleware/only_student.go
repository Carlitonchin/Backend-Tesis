package middleware

import (
	"os"

	"github.com/Carlitonchin/Backend-Tesis/handler/handler_utils"
	"github.com/Carlitonchin/Backend-Tesis/model/apperrors"
	"github.com/gin-gonic/gin"
)

func OnlyStudent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := handler_utils.GetUser(ctx)
		if err != nil {
			handler_utils.SendErrorResponse(ctx, err)
			ctx.Abort()
			return
		}

		if user.Role.Name != os.Getenv("ROLE_DEFAULT_STUDENT") {
			type_error := apperrors.Authorization
			message := "No tiene permisos para acceder a este recurso"

			e := apperrors.NewError(type_error, message)
			ctx.JSON(e.Status(), gin.H{
				"error": e,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
