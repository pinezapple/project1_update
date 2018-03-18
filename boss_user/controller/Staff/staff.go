package Staff

import (
	"net/http"
	"project1_update/boss_user/controller"

	"project1_update/boss_user/microservice/staffmana"
	"strconv"

	"github.com/labstack/echo"
)

func AddStaff(c echo.Context) (erro error) {

	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	name := c.FormValue("name")
	passwd := c.FormValue("passwd")
	position, _ := strconv.Atoi(c.FormValue("pos"))

	ok, err := staffmana.AddStaff(name, passwd, int64(position))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}
	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelStaff(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	ok, err := staffmana.DelStaff(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}
	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllStaff(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := staffmana.SelectAllStaff()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectStaffById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := staffmana.SelectStaffById(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
