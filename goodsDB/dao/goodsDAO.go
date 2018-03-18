package dao

import (
	"database/sql"
	"log"
	goodsModel "project1_update/others/models/goodsmod"
	"strconv"
)

type Igoods interface {
	AddGoods(goods *goodsModel.GoodsMod) (bool, error)
	DelGoods(id int64) (bool, error)
	UpdateGoods(id int64, new *goodsModel.GoodsMod) (bool, error)
	GoodsExist(id int64) (bool, error)
	SelectAll() ([]*goodsModel.GoodsMod, error)
	SelectByID(id int64) (*goodsModel.GoodsMod, bool, error)
	SelectByName(name string) ([]*goodsModel.GoodsMod, bool, error)
	SelectByPrice(price int64) ([]*goodsModel.GoodsMod, bool, error)
}

type goodsDao struct {
}

func (c goodsDao) AddGoods(goods *goodsModel.GoodsMod) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return false, err1
	}
	s := "'" + goods.Name + "'"
	s1 := "'" + strconv.FormatInt(goods.Quantity, 10) + "'"
	s2 := "'" + strconv.FormatInt(goods.Price, 10) + "'"
	_, err2 := db.Exec("insert into goods_table (id,name,quantity,price) values NULL," + s + "," + s1 + "," + s2 + ")")
	if err2 != nil {
		log.Fatalf("Can't add goods %v", err2)
		return false, err2
	}

	return true, nil
}

func (c goodsDao) DelGoods(id int64) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from goods_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete goods id number %d", id)
		return false, err2
	}

	return true, nil
}

func (c goodsDao) UpdateGoods(id int64, new *goodsModel.GoodsMod) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return false, err1
	}

	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	s2 := "'" + new.Name + "'"
	s3 := "'" + strconv.FormatInt(new.Quantity, 10) + "'"
	s4 := "'" + strconv.FormatInt(new.Price, 10) + "'"

	_, err2 := db.Exec("update goods_table set name=" + s2 + ", quantity=" + s3 + ", price=" + s4 + "where id = " + s1)
	if err2 != nil {
		log.Fatalf("Can't update goods_db %v", err2)
		return false, err2
	}
	return true, nil
}

func (c goodsDao) GoodsExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from goods_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for goods", err)
		return false, err
	}
	id1 := 0
	for rows.Next() {
		err := rows.Scan(id1)
		if err != nil {
			log.Fatal("Can't scan for goods", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c goodsDao) SelectAll() (r []*goodsModel.GoodsMod, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from goods_table")

	if err != nil {
		log.Fatalf("Can't search for all goods", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Quantity, &r[i].Price)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for goods", err)
			return nil, err
		}
	}
	return r, nil
}

func (c goodsDao) SelectByID(id int64) (r *goodsModel.GoodsMod, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to goods Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from goods_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for goods by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Name, &r.Quantity, &r.Price)
		if err != nil {
			log.Fatalf("Can't scan for goods %v", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c goodsDao) SelectByName(name string) (r []*goodsModel.GoodsMod, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return nil, false, err
	}
	s := "'" + name + "'"
	rows, err := db.Query("select * from goods_table where name =" + s)

	if err != nil {
		log.Fatalf("Can't search for goods by name", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Quantity, &r[i].Price)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for goods", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

func (c goodsDao) SelectByPrice(price int64) (r []*goodsModel.GoodsMod, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use goods_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Goods Database %v", err1)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(price, 10) + "'"
	rows, err := db.Query("select * from goods_table where price =" + s)

	if err != nil {
		log.Fatalf("Can't search for goods by price", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Quantity, &r[i].Price)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for goods", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

var GoodsDao Igoods = goodsDao{}
