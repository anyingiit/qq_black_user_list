package e

import (
	"github.com/gin-gonic/gin"
)

func ResponseForJson(ctx *gin.Context, data interface{}, err *Error) {
	if err != nil {
		if err.MoreInfo != "" {
			ctx.JSON(err.Status, &FaildResponse{
				Code:     err.Code,
				Message:  err.Message,
				MoreInfo: err.MoreInfo,
			})
			ctx.Abort()
			return
		}
		ctx.JSON(err.Status, &Base{
			Code:    err.Code,
			Message: err.Message,
		})
		ctx.Abort()
		return
	} else {
		if data != nil {
			//fmt.Println("in")
			ctx.JSON(OK.Status, &SuccessResponse{
				Base: &Base{
					Code:    OK.Code,
					Message: OK.Message,
				},
				Data: data,
			})
			ctx.Abort()
			return
		}
		ctx.JSON(OK.Status, &Base{
			Code:    OK.Code,
			Message: OK.Message,
		})
		ctx.Abort()
		return
	}
}
