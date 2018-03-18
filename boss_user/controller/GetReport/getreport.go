package GetReport

import (
	"net/http"
	"project1_update/boss_user/controller"
	"project1_update/boss_user/microservice/goods_reporter"
	"strconv"

	"github.com/labstack/echo"
)

func IfGoodExist(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	ok, err := goods_reporter.IfGoodExist(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllGoods(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := goods_reporter.SelectAllGoods()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectGoodsByID(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	payload, err := goods_reporter.SelectGoodsById(int64(id))

	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectGoodsByName(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	name := c.FormValue("name")
	payload, err := goods_reporter.SelectGoodsByName(name)

	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectGoodsByPrice(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	price, _ := strconv.Atoi(c.FormValue("price"))
	payload, err := goods_reporter.SelectGoodsByPrice(int64(price))

	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
