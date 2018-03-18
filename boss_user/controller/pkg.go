package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// ResponseError ...
type ResponseError struct {
	Code    int
	Message string
}

// Response ...
type Response struct {
	Error ResponseError
	Data  interface{}
}

// SetData ...
func (c *Response) SetData(_dat interface{}) {
	c.Data = _dat
}

// SetCodeMessage ...
func (c *Response) SetCodeMessage(code int, message string) {
	c.Error.Code = code
	c.Error.Message = string(message)
}

// HandlerTemplate template for handler
func HandlerTemplate(c echo.Context, cmd func(c echo.Context) (*Response, error)) (erro error) {
	erro = nil
	response := &Response{}

	defer func() {
		if erro == nil {
			erro = c.JSON(http.StatusOK, response)
		}
	}()

	response, erro = cmd(c)
	return
}
