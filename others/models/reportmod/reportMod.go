package models

import (
	"time"
)

type DonSell struct {
	Id                int64
	Seller_name       string
	Cus_name          string
	Goods_list        []byte
	Goods_list_amount []byte
	Total_money       int64
	Created_at        time.Time
	Updated_at        time.Time
}

type DonImport struct {
	Id                int64
	Created_by        string
	Goods_list        []byte
	Goods_list_amount []byte
	Total_money       int64
	Created_at        string
	Updated_at        string
}

type DonSpend struct {
	Id          int64
	Created_by  string
	Reason      string
	Total_money int64
	Created_at  string
	Updated_at  string
}

type DonExport struct {
	Id                int64
	Created_by        string
	Reason            string
	Goods_list        []byte
	Goods_list_amount []byte
	Total_money       int64
	Created_at        string
	Updated_at        string
}

type HoursByDay struct {
	Id      int64
	StaffId int64
	In      int64
	Out     int64
}
