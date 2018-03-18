package Spend

import (
	"net/http"
	"project1_update/boss_user/controller"
	"project1_update/boss_user/microservice/spender"
	"strconv"

	"github.com/labstack/echo"
)

func AddSpend(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	created_by := c.FormValue("createdby")
	reason := c.FormValue("reason")
	total_money, _ := strconv.Atoi(c.FormValue("total"))

	ok, err := spender.AddDSp(created_by, reason, int64(total_money))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelSpend(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	ok, err := spender.DelDSp(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllSpend(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := spender.SelectAllDSp()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectSpendById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := spender.SelectByIDDSp(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectSpendByCreatedAt(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	from_time, _ := strconv.Atoi(c.FormValue("fromtime"))
	to_time, _ := strconv.Atoi(c.FormValue("totime"))

	payload, err := spender.SelectByCreatedAtDSp(int64(from_time), int64(to_time))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
