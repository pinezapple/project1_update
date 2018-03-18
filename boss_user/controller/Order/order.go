package Order

import (
	"net/http"
	"project1_update/boss_user/controller"

	"project1_update/boss_user/microservice/orderer"
	"strconv"

	"github.com/labstack/echo"
)

func AddOrder(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	seller_name := c.FormValue("seller")
	cus_name := c.FormValue("cus")
	goods_list := c.FormValue("goodslist")
	goods_list_amount := c.FormValue("amount")

	ok, err := orderer.AddDS(seller_name, cus_name, []byte(goods_list), []byte(goods_list_amount))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func DelOrder(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))
	ok, err := orderer.DelDS(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(ok)
	return
}

func SelectAllOrder(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	payload, err := orderer.SelectAllDS()
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectOrderById(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	id, _ := strconv.Atoi(c.FormValue("id"))

	payload, err := orderer.SelectByIDDS(int64(id))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}

func SelectOrderByCreatedAt(c echo.Context) (erro error) {
	response := controller.Response{}
	defer func() {
		erro = c.JSON(http.StatusOK, response)
	}()

	from_time, _ := strconv.Atoi(c.FormValue("fromtime"))
	to_time, _ := strconv.Atoi(c.FormValue("totime"))

	payload, err := orderer.SelectByCreatedAtDS(int64(from_time), int64(to_time))
	if err != nil {
		response.SetCodeMessage(http.StatusBadGateway, "")
	}

	response.SetCodeMessage(http.StatusOK, "")
	response.SetData(payload)
	return
}
