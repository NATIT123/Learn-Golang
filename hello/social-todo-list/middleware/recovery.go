package middleware

import (
	"errors"
	"fmt"
	"main/common"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				// Xử lý nếu err là *common.AppError
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return // không nên panic lại
				}

				// Dùng type switch để ép kiểu an toàn
				var realErr error
				switch e := err.(type) {
				case error:
					realErr = e
				case string:
					realErr = errors.New(e)
				default:
					realErr = fmt.Errorf("unknown error: %v", e)
				}

				appErr := common.ErrInternal(realErr)
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// không nên panic lại ở đây
			}
		}()
		c.Next()
	}
}
