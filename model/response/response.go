package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 500
	SUCCESS = 200
	UNLOGIN = 416
)

type Response struct {
	ResultCode int         `json:"resultCode"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"message"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {

	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "SUCCESS", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func UnLogin(data interface{}, c *gin.Context) {
	Result(UNLOGIN, data, "not login", c)
}
