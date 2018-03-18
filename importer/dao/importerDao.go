package dao

import (
	"database/sql"
	"log"

	DonModel "project1_update/others/models/reportmod"
	"strconv"
	"time"
)

type IImporter interface {
	AddDI(don *DonModel.DonImport) (bool, error)
	DelDI(id int64) (bool, []byte, []byte, error)
	UpdateDI(id int64, new *DonModel.DonImport) (bool, error)
	DIExist(id int64) (bool, error)
	SelectAll() ([]*DonModel.DonImport, error)
	SelectByID(id int64) (*DonModel.DonImport, bool, error)
	SelectByCreatedAt(from_time int64, to_time int64) ([]*DonModel.DonImport, bool, error)
}

type importerDao struct {
}

func (c importerDao) AddDI(don *DonModel.DonImport) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Import Database %v", err1)
		return false, err1
	}

	s1 := "'" + don.Created_by + "'"
	s2 := "'" + string(don.Goods_list) + "'"
	s3 := "'" + string(don.Goods_list_amount) + "'"
	s4 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s5 := "'" + time.Now().String() + "'"
	s6 := "'" + time.Now().String() + "'"
	_, err2 := db.Exec("insert into import_table (id,created_by,goods_list,goods_list_amount,total_money,created_at,updated_at) values NULL," + s1 + "," + s2 + "," + s3 + "," + s4 + "," + s5 + "," + s6 + ")")
	if err2 != nil {
		log.Fatalf("Can't add to Import Table %v", err2)
		return false, err2
	}

	return true, nil
}

func (c importerDao) DelDI(id int64) (ok bool, goods_list []byte, goods_list_amount []byte, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, nil, nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Import Database %v", err1)
		return false, nil, nil, err1
	}

	//del good_list
	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select goods_list,goods_list_amount from import_table where id=" + s1)

	if err != nil {
		log.Fatalf("Can't search for Don Import", err)
		return false, nil, nil, err
	}

	for rows.Next() {
		err := rows.Scan(&goods_list, &goods_list_amount)
		if err != nil {
			log.Fatal("Can't scan for Don Import", err)
			return false, nil, nil, err
		}
		/*
			err1 := os.Remove(goods_list)
			if err1 != nil {
				log.Fatalf("Can't remove the old import good list", err)
				return false, err
			}*/
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from import_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete from Import tablev id number %d", id)
		return false, nil, nil, err2
	}

	return true, goods_list, goods_list_amount, nil
}

func (c importerDao) UpdateDI(id int64, don *DonModel.DonImport) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to import Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(don.Id, 10) + "'"
	s1 := "'" + don.Created_by + "'"
	s2 := "'" + string(don.Goods_list) + "'"
	s3 := "'" + string(don.Goods_list_amount) + "'"
	s4 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s5 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("update import_table set create_by=" + s1 + ", goods_list=" + s2 + ",goods_list_amount=" + s3 + ", total_money=" + s4 + ",updated_at= " + s5 + "where id = " + s)
	if err2 != nil {
		log.Fatalf("Can't update import table %v", err2)
		return false, err2
	}
	return true, nil
}

func (c importerDao) DIExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Import Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from import_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for all Import", err)
		return false, err
	}

	id1 := 0
	for rows.Next() {
		err := rows.Scan(&id1)
		if err != nil {
			log.Fatal("Can't scan for Don Import", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c importerDao) SelectAll() (r []*DonModel.DonImport, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err != nil {
		log.Fatalf("Can't connect to Importer Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from import_table")

	if err != nil {
		log.Fatalf("Can't search for all import", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for import", err)
			return nil, err
		}
	}
	return r, nil
}

func (c importerDao) SelectByID(id int64) (r *DonModel.DonImport, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Import Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from import_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for import by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Created_by, &r.Goods_list, &r.Goods_list_amount, &r.Total_money, &r.Created_at, &r.Updated_at)
		if err != nil {
			log.Fatalf("Can't scan for Don Report %v", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c importerDao) SelectByCreatedAt(from_time int64, to_time int64) (r []*DonModel.DonImport, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use importer_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Import Database %v", err1)
		return nil, false, err
	}
	tin := time.Unix(from_time, 0)
	tin1 := tin.Format(time.UnixDate)

	s1 := "'" + tin1 + "'"

	tout := time.Unix(to_time, 0)
	tout1 := tout.Format(time.UnixDate)

	s2 := "'" + tout1 + "'"
	rows, err := db.Query("select * from import_table where (created_at <=" + s2 + ") and (" + s1 + "<= created_at)")

	if err != nil {
		log.Fatalf("Can't search for import by created_at", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for import", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

var ImporterDao IImporter = importerDao{}
