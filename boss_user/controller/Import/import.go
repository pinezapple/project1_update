package Import

import (
	"net/http"
	"project1_update/boss_user/controller"

	"project1_update/boss_user/microservice/importer"
	"strconv"

	"github.com/labstack/echo"
)

func AddImport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	created_by := c.FormValue("createdby")
	goods_list := c.FormValue("goodslist")
	goods_list_amount := c.FormValue("amount")

	ok, err := importer.AddDI(created_by, []byte(goods_list), []byte(goods_list_amount))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelImport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	ok, err := importer.DelDI(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllImport(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := importer.SelectAllDI()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectImportById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := importer.SelectByIDDI(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectImportByCreatedAt(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	from_time, _ := strconv.Atoi(c.FormValue("fromtime"))
	to_time, _ := strconv.Atoi(c.FormValue("totime"))

	payload, err := importer.SelectByCreatedAtDI(int64(from_time), int64(to_time))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
