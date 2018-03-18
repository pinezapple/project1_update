package Authen

import (
	"net/http"
	"project1_update/boss_user/controller"
	"project1_update/boss_user/microservice/authen"
	"strconv"

	"github.com/labstack/echo"
)

func Signin(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	success, _, err := ProcessSignin(c)
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}
	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(success)
	return
}

func ProcessSignin(c echo.Context) (ok bool, code int, err error) {

	staff_id, _ := strconv.Atoi(c.FormValue("id"))
	passwd := c.FormValue("passwd")
	position, _ := strconv.Atoi(c.FormValue("pos"))

	ok, err = authen.AuthenAndClassify(int64(staff_id), passwd, int64(position))
	if err != nil {
		return false, http.StatusBadGateway, err
	}
	return ok, http.StatusOK, nil
}
