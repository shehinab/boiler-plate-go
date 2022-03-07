package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Response -
type Response struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

//Error -  common json error
type Error struct {
	Status      bool   `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"error"`
	trace       string
}
type errorString struct {
	s string
}

//SuccessResponse -
func SuccessResponse(c *gin.Context, data interface{}) Response {
	response := Response{
		Code:   http.StatusOK,
		Status: true,
		Data:   data,
	}
	return response
}

func (e *errorString) Error() string {
	return e.s
}

//CustomError for error handling
func CustomError(text string) error {
	return &errorString{text}
}

//InternalServerError - internal server error
func InternalServerError(c *gin.Context, e error) Error {

	errMsg := Error{
		Status:      false,
		Code:        http.StatusInternalServerError,
		Description: e.Error(),
	}

	return errMsg
}

//BadDataRequest - validator response
func BadDataRequest(c *gin.Context, e error) Error {

	errMsg := Error{
		Status:      false,
		Code:        http.StatusBadRequest,
		Description: e.Error(),
	}

	return errMsg
}

//PreConditionFailed - validator response
func PreConditionFailed(c *gin.Context, e error) Error {

	errMsg := Error{
		Status:      false,
		Code:        http.StatusPreconditionFailed,
		Description: e.Error(),
	}

	return errMsg
}

//Unauthorised - validator response
func Unauthorised(c *gin.Context, e error) Error {

	errMsg := Error{
		Status:      false,
		Code:        http.StatusUnauthorized,
		Description: e.Error(),
	}

	return errMsg
}

//Forbidden - validator response
func Forbidden(c *gin.Context, e error) Error {

	errMsg := Error{
		Status:      false,
		Code:        http.StatusForbidden,
		Description: e.Error(),
	}

	return errMsg
}
