package main

import (
	"project1_update/boss_user/controller/Authen"
	"project1_update/boss_user/controller/Customer"
	"project1_update/boss_user/controller/Export"
	"project1_update/boss_user/controller/GetReport"
	"project1_update/boss_user/controller/Import"
	"project1_update/boss_user/controller/Order"
	"project1_update/boss_user/controller/Spend"
	"project1_update/boss_user/controller/Staff"

	"github.com/labstack/echo"
)

func handlers(router *echo.Echo) {

	router.POST("/authen", Authen.Signin)

	cusgrp := router.Group("/cus")
	cusgrp.POST("/add", Customer.AddCustomer)
	cusgrp.POST("/del", Customer.DelCus)
	cusgrp.POST("selectall", Customer.SelectAllCus)
	cusgrp.POST("/selectbyid", Customer.SelectCusById)
	cusgrp.POST("/selectbyname", Customer.SelectCusByName)
	cusgrp.POST("/selectbyphonenum", Customer.SelectCusByPhoneNum)

	staffgrp := router.Group("/staff")
	staffgrp.POST("/add", Staff.AddStaff)
	staffgrp.POST("/del", Staff.DelStaff)
	staffgrp.POST("/selectall", Staff.SelectAllStaff)
	staffgrp.POST("/selectbyid", Staff.SelectStaffById)

	goodsrp := router.Group("/goods")
	goodsrp.POST("/check", GetReport.IfGoodExist)
	goodsrp.POST("/selectall", GetReport.SelectAllGoods)
	goodsrp.POST("/selectbyid", GetReport.SelectGoodsByID)
	goodsrp.POST("/selectbyname", GetReport.SelectGoodsByName)
	goodsrp.POST("/selectbyprice", GetReport.SelectGoodsByPrice)

	exportgrp := router.Group("/exp")
	exportgrp.POST("/add", Export.AddExport)
	exportgrp.POST("/del", Export.DelExport)
	exportgrp.POST("/selectall", Export.SelectAllExport)
	exportgrp.POST("/selectbyid", Export.SelectExportById)
	exportgrp.POST("/selectbycreatedat", Export.SelectExportByCreatedAt)

	importgrp := router.Group("/imp")
	importgrp.POST("/add", Import.AddImport)
	importgrp.POST("/del", Import.DelImport)
	importgrp.POST("/selectall", Import.SelectAllImport)
	importgrp.POST("/selectbyid", Import.SelectImportById)
	importgrp.POST("/selectbycreatedat", Import.SelectImportByCreatedAt)

	ordergrp := router.Group("/ord")
	ordergrp.POST("/add", Order.AddOrder)
	ordergrp.POST("/del", Order.DelOrder)
	ordergrp.POST("/selectall", Order.SelectAllOrder)
	ordergrp.POST("/selectbyid", Order.SelectOrderById)
	ordergrp.POST("/selectbycreatedat", Order.SelectOrderByCreatedAt)

	spendgrp := router.Group("/spd")
	spendgrp.POST("/add", Spend.AddSpend)
	spendgrp.POST("/del", Spend.DelSpend)
	spendgrp.POST("/selectall", Spend.SelectAllSpend)
	spendgrp.POST("/selectbyid", Spend.SelectSpendById)
	spendgrp.POST("/selectbycreatedat", Spend.SelectSpendByCreatedAt)

}

func main() {
	e := echo.New()
	handlers(e)
	e.Start(":8080")
}
