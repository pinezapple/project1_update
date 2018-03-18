package dao

import (
	"database/sql"
	"log"
	"os"
	DonModel "project1_update/others/models/reportmod"
	"strconv"
	"time"
)

type ISpender interface {
	AddDSp(don *DonModel.DonSpend) (bool, error)
	DelDSp(id int64) (bool, error)
	UpdateDSp(id int64, new *DonModel.DonSpend) (bool, error)
	DSpExist(id int64) (bool, error)
	SelectAll() ([]*DonModel.DonSpend, error)
	SelectByID(id int64) (*DonModel.DonSpend, bool, error)
	SelectByCreatedAt(from_time int64, to_time int64) ([]*DonModel.DonSpend, bool, error)
}

type spenderDao struct {
}

func (c spenderDao) AddDSp(don *DonModel.DonSpend) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return false, err1
	}

	s1 := "'" + don.Created_by + "'"
	s2 := "'" + don.Reason + "'"
	s3 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s4 := "'" + time.Now().String() + "'"
	s5 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("insert into spend_table (id,created_by,reason,total_money,created_at,updated_at) values NULL," + s1 + "," + s2 + "," + s3 + "," + s4 + "," + s5 + ")")
	if err2 != nil {
		log.Fatalf("Can't add Don Spend %v", err2)
		return false, err2
	}

	return true, nil
}

func (c spenderDao) DelDSp(id int64) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return false, err1
	}

	//del good_list
	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select good_list from spend_table where id=" + s1)

	if err != nil {
		log.Fatalf("Can't search for Don Spend", err)
		return false, err
	}

	var goods_list string
	for rows.Next() {
		err := rows.Scan(&goods_list)
		if err != nil {
			log.Fatal("Can't scan for Don Spend", err)
			return false, err
		}
		err1 := os.Remove(goods_list)
		if err1 != nil {
			log.Fatalf("Can't remove the old spend good list", err)
			return false, err
		}
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from spend_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete Don Spend id number %d", id)
		return false, err2
	}

	return true, nil
}

func (c spenderDao) UpdateDSp(id int64, don *DonModel.DonSpend) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(don.Id, 10) + "'"
	s1 := "'" + don.Created_by + "'"
	s2 := "'" + don.Reason + "'"
	s3 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s5 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("update spend_table set created_by=" + s1 + ", reason=" + s2 + ", total_money=" + s3 + ",updated_at= " + s5 + "where id = " + s)
	if err2 != nil {
		log.Fatalf("Can't update spend_db %v", err2)
		return false, err2
	}
	return true, nil
}

func (c spenderDao) DSpExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from spend_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for Don Spend", err)
		return false, err
	}

	id1 := 0
	for rows.Next() {
		err := rows.Scan(id1)
		if err != nil {
			log.Fatal("Can't scan for Don Spend", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c spenderDao) SelectAll() (r []*DonModel.DonSpend, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from spend_table")

	if err != nil {
		log.Fatalf("Can't search for all spend", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Reason, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for spend", err)
			return nil, err
		}
	}
	return r, nil
}

func (c spenderDao) SelectByID(id int64) (r *DonModel.DonSpend, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from spend_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for spend by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Created_by, &r.Reason, &r.Total_money, &r.Created_at, &r.Updated_at)
		if err != nil {
			log.Fatal("Can't scan for spend by id", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c spenderDao) SelectByCreatedAt(from_time int64, to_time int64) (r []*DonModel.DonSpend, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use spender_db")
	if err1 != nil {
		log.Fatalf("Can't connect to spend Database %v", err1)
		return nil, false, err
	}
	tin := time.Unix(from_time, 0)
	tin1 := tin.Format(time.UnixDate)

	s1 := "'" + tin1 + "'"

	tout := time.Unix(to_time, 0)
	tout1 := tout.Format(time.UnixDate)

	s2 := "'" + tout1 + "'"

	rows, err := db.Query("select * from spend_table where (created_at <=" + s2 + ") and (" + s1 + "<= created_at)")

	if err != nil {
		log.Fatalf("Can't search for spend by created_at", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Reason, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for sell", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

var SpenderDao ISpender = spenderDao{}
