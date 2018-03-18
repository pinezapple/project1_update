package dao

import (
	"database/sql"
	"log"
	DonModel "project1_update/others/models/reportmod"
	"strconv"
	"time"
)

type IOrderer interface {
	AddDS(don *DonModel.DonSell) (bool, error)
	DelDS(id int64) (bool, int64, string, []byte, []byte, error)
	UpdateDS(id int64, new *DonModel.DonSell) (bool, error)
	DSExist(id int64) (bool, error)
	SelectAll() ([]*DonModel.DonSell, error)
	SelectByID(id int64) (*DonModel.DonSell, bool, error)
	SelectByCreatedAt(from_time int64, to_time int64) ([]*DonModel.DonSell, bool, error)
}

type ordererDao struct {
}

func (c ordererDao) AddDS(don *DonModel.DonSell) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err1)
		return false, err1
	}

	s1 := "'" + don.Seller_name + "'"
	s2 := "'" + don.Cus_name + "'"
	s3 := "'" + string(don.Goods_list) + "'"
	s4 := "'" + string(don.Goods_list_amount) + "'"
	s5 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s6 := "'" + time.Now().String() + "'"
	s7 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("insert into order_table (id,seller_name,cus_name,goods_list,goods_list_amount,total_money,created_at,updated_at) values NULL," + s1 + "," + s2 + "," + s3 + "," + s4 + "," + s5 + "," + s6 + "," + s7 + ")")
	if err2 != nil {
		log.Fatalf("Can't add Don Sell %v", err2)
		return false, err2
	}

	return true, nil
}

func (c ordererDao) DelDS(id int64) (ok bool, total_money int64, cusname string, goods_list []byte, goods_list_amount []byte, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, 0, "", nil, nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err1)
		return false, 0, "", nil, nil, err1
	}

	//del good_list
	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select cus_name,total_money, goods_list,goods_list_amount from order_table where id=" + s1)

	if err != nil {
		log.Fatalf("Can't search for Don Order", err)
		return false, 0, "", nil, nil, err
	}

	for rows.Next() {
		err := rows.Scan(&cusname, &total_money, &goods_list, &goods_list_amount)
		if err != nil {
			log.Fatal("Can't scan for Don Order", err)
			return false, 0, "", nil, nil, err
		}
		/*
			err1 := os.Remove(goods_list)
			if err1 != nil {
				log.Fatalf("Can't remove the old order good list", err)
				return false, nil, false, err
			} */
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from order_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete Don Sell id number %d", id)
		return false, 0, "", nil, nil, err2
	}

	return true, total_money, cusname, goods_list, goods_list_amount, nil
}

func (c ordererDao) UpdateDS(id int64, don *DonModel.DonSell) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(don.Id, 10) + "'"
	s1 := "'" + don.Seller_name + "'"
	s2 := "'" + don.Cus_name + "'"
	s3 := "'" + string(don.Goods_list) + "'"
	s4 := "'" + string(don.Goods_list_amount) + "'"
	s5 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s6 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("update order_table set seller_name=" + s1 + ", cus_name=" + s2 + ", goods_list=" + s3 + ", goods_list_amount=" + s4 + ",total_money=" + s5 + ",updated_at= " + s6 + "where id = " + s)
	if err2 != nil {
		log.Fatalf("Can't update sell_db %v", err2)
		return false, err2
	}
	return true, nil
}

func (c ordererDao) DSExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from order_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for Don Sell", err)
		return false, err
	}

	id1 := 0
	for rows.Next() {
		err := rows.Scan(id1)
		if err != nil {
			log.Fatal("Can't scan for Don Sell", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c ordererDao) SelectAll() (r []*DonModel.DonSell, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err != nil {
		log.Fatalf("Can't connect to orderer Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from order_table")

	if err != nil {
		log.Fatalf("Can't search for all sell", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Seller_name, &r[i].Cus_name, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for sell", err)
			return nil, err
		}
	}
	return r, nil
}

func (c ordererDao) SelectByID(id int64) (r *DonModel.DonSell, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from order_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for sell by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Seller_name, &r.Cus_name, &r.Goods_list, &r.Goods_list_amount, &r.Total_money, &r.Created_at, &r.Updated_at)
		if err != nil {
			log.Fatal("Can't scan for sell by id", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c ordererDao) SelectByCreatedAt(from_time int64, to_time int64) (r []*DonModel.DonSell, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use orderer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to order Database %v", err1)
		return nil, false, err
	}

	tin := time.Unix(from_time, 0)
	tin1 := tin.Format(time.UnixDate)

	s1 := "'" + tin1 + "'"

	tout := time.Unix(to_time, 0)
	tout1 := tout.Format(time.UnixDate)

	s2 := "'" + tout1 + "'"

	rows, err := db.Query("select * from order_table where (created_at <=" + s2 + ") and (" + s1 + "<= created_at)")

	if err != nil {
		log.Fatalf("Can't search for sell by created_at", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Seller_name, &r[i].Cus_name, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for sell", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

var OrdererDao IOrderer = ordererDao{}
