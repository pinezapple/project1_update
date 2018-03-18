package Customer

import (
	"net/http"
	"project1_update/boss_user/controller"
	"project1_update/boss_user/microservice/cusmana"
	"strconv"

	"github.com/labstack/echo"
)

func AddCustomer(c echo.Context) (erro error) {

	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	name := c.FormValue("name")
	phonenum := c.FormValue("phonenum")
	balance, _ := strconv.Atoi(c.FormValue("balance"))

	ok, err := cusmana.AddCus(name, phonenum, int64(balance))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}
	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelCus(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	ok, err := cusmana.DelCus(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}
	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllCus(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := cusmana.SelectAllCus()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectCusById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := cusmana.SelectCusById(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectCusByName(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	name := c.FormValue("name")

	payload, err := cusmana.SelectCusByName(name)
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectCusByPhoneNum(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	phonenum := c.FormValue("phonenum")

	payload, err := cusmana.SelectCusByPhoneNum(phonenum)
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
