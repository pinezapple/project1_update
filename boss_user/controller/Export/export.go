package Export

import (
	"net/http"
	"project1_update/boss_user/controller"
	"project1_update/boss_user/microservice/exporter"
	"strconv"

	"github.com/labstack/echo"
)

func AddExport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	created_by := c.FormValue("createdby")
	reason := c.FormValue("reason")
	goods_list := c.FormValue("goodslist")
	goods_list_amount := c.FormValue("amount")

	ok, err := exporter.AddDE(created_by, reason, []byte(goods_list), []byte(goods_list_amount))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelExport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	ok, err := exporter.DelDE(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllExport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := exporter.SelectAllDE()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectExportById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := exporter.SelectByIDDE(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectExportByCreatedAt(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	from_time, _ := strconv.Atoi(c.FormValue("fromtime"))
	to_time, _ := strconv.Atoi(c.FormValue("totime"))

	payload, err := exporter.SelectByCreatedAtDE(int64(from_time), int64(to_time))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
