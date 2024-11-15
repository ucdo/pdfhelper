package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RespType int

const (
	done RespType = 200
	fail RespType = 500
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    int(done),
		Message: "success",
		Data:    data,
	})
}

// Fail 失败
func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    int(fail),
		Message: message,
		Data:    struct{}{},
	})
}

func Error(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		Fail(c, err.Error())
		return
	}

	Fail(c, err.Error())
}
