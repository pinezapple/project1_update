package dao

import (
	"database/sql"
	"log"
	cusModel "project1_update/others/models/customermod"
	"strconv"
)

type ICus interface {
	AddCus(cus *cusModel.CustomerMod) (bool, error)
	DelCus(id int64) (bool, error)
	UpdateCus(id int64, new *cusModel.CustomerMod) (bool, error)
	CusExist(id int64) (bool, error)
	SelectAll() ([]*cusModel.CustomerMod, error)
	SelectByID(id int64) (*cusModel.CustomerMod, bool, error)
	SelectByName(name string) ([]*cusModel.CustomerMod, bool, error)
	SelectByPhoneNum(phonenum string) (*cusModel.CustomerMod, bool, error)
	AddToBalance(id int64, amount int64) (bool, error)
}

type cusDao struct {
}

func (c cusDao) AddCus(cus *cusModel.CustomerMod) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return false, err1
	}
	s := "'" + cus.Phonenum + "'"
	s1 := "'" + cus.Name + "'"
	s2 := "'" + strconv.FormatInt(cus.Balance, 10) + "'"
	s3 := "'" + strconv.FormatInt(int64(cus.Level), 10) + "'"
	_, err2 := db.Exec("insert into cus_table (id,name,phonenum,balance,level) values NULL," + s1 + "," + s + "," + s2 + "," + s3 + ")")
	if err2 != nil {
		log.Fatalf("Can't add customer %v", err2)
		return false, err2
	}

	return true, nil
}

func (c cusDao) DelCus(id int64) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from cus_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete customer id number %d", id)
		return false, err2
	}

	return true, nil
}

func (c cusDao) UpdateCus(id int64, new *cusModel.CustomerMod) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return false, err1
	}

	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	s2 := "'" + new.Name + "'"
	s3 := "'" + new.Phonenum + "'"
	s4 := "'" + strconv.FormatInt(new.Balance, 10) + "'"
	s5 := "'" + strconv.FormatInt(int64(new.Level), 10) + "'"
	_, err2 := db.Exec("update cus_table set name=" + s2 + ", phonenum=" + s3 + ", balance=" + s4 + ", level=" + s5 + "where id = " + s1)
	if err2 != nil {
		log.Fatalf("Can't update cus_db %v", err2)
		return false, err2
	}
	return true, nil
}

func (c cusDao) CusExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from cus_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for customers", err)
		return false, err
	}
	id1 := 0
	for rows.Next() {
		err := rows.Scan(id1)
		if err != nil {
			log.Fatal("Can't scan for customer", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c cusDao) SelectAll() (r []*cusModel.CustomerMod, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from cus_table")

	if err != nil {
		log.Fatalf("Can't search for all customers", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Phonenum, &r[i].Balance, &r[i].Level)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for customer", err)
			return nil, err
		}
	}
	return r, nil
}

func (c cusDao) SelectByID(id int64) (r *cusModel.CustomerMod, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from cus_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for customers by id", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Name, &r.Phonenum, &r.Balance, &r.Level)
		if err != nil {
			log.Fatal("Can't scan for customer", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c cusDao) SelectByName(name string) (r []*cusModel.CustomerMod, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return nil, false, err
	}
	s := "'" + name + "'"
	rows, err := db.Query("select * from cus_table where name =" + s)

	if err != nil {
		log.Fatalf("Can't search for customers by name", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Phonenum, &r[i].Balance, &r[i].Level)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for customer", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

func (c cusDao) SelectByPhoneNum(phonenum string) (r *cusModel.CustomerMod, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return nil, false, err
	}
	s := "'" + phonenum + "'"
	rows, err := db.Query("select * from cus_table where phonenum =" + s)

	if err != nil {
		log.Fatalf("Can't search for customers by phone number", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Name, &r.Phonenum, &r.Balance, &r.Level)
		if err != nil {
			log.Fatal("Can't scan for customer", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

func (c cusDao) AddToBalance(id int64, amount int64) (ok bool, err error) {
	var balance int64
	var level int32
	amountperlevel := [6]int64{500, 1000, 2000, 3500, 4000, 5000}
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use cus_db")
	if err != nil {
		log.Fatalf("Can't connect to Customer Database %v", err1)
		return false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select balance,level from cus_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for customers by id", err)
		return false, err
	}
	for rows.Next() {
		err := rows.Scan(balance, level)
		if err != nil {
			log.Fatal("Can't scan for customer's balance ", err)
			return false, err
		}
		if amount+balance >= amountperlevel[level]*1000 {
			balance := balance + amount
			level := level + 1
			s1 := "'" + strconv.FormatInt(balance, 10) + "'"
			s2 := "'" + strconv.FormatInt(int64(level), 10) + "'"
			_, err1 := db.Exec("update cus_table set balance=" + s1 + ", level=" + s2 + "where id = " + s)
			if err1 != nil {
				log.Fatalf("can't add to cus_balance and update cus level %v", err1)
				return false, err1
			}
		} else {
			balance := balance + amount
			s1 := "'" + strconv.FormatInt(balance, 10) + "'"
			_, err1 := db.Exec("update cus_table set balance=" + s1 + " where id = " + s)
			if err1 != nil {
				log.Fatalf("can't add to cus_balance and update cus level %v", err1)
				return false, err1
			}
		}
	}
	return true, nil
}

var CusDao ICus = cusDao{}
