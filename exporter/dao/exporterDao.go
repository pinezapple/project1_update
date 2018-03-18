package dao

import (
	"database/sql"
	"log"
	DonModel "project1_update/others/models/reportmod"
	"strconv"
	"time"
)

type IExporter interface {
	AddDE(don *DonModel.DonExport) (bool, error)
	DelDE(id int64) (bool, []byte, []byte, error)
	UpdateDE(id int64, new *DonModel.DonExport) (bool, error)
	DEExist(id int64) (bool, error)
	SelectAll() ([]*DonModel.DonExport, error)
	SelectByID(id int64) (*DonModel.DonExport, bool, error)
	SelectByCreatedAt(from_time int64, to_time int64) ([]*DonModel.DonExport, bool, error)
}

type exporterDao struct {
}

func (c exporterDao) AddDE(don *DonModel.DonExport) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return false, err1
	}

	s1 := "'" + don.Created_by + "'"
	s2 := "'" + don.Reason + "'"
	s3 := "'" + string(don.Goods_list) + "'"
	s4 := "'" + string(don.Goods_list_amount) + "'"
	s5 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s6 := "'" + time.Now().String() + "'"
	s7 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("insert into export_table (id,created_by,reason,goods_list,goods_list_amount,total_money,created_at,updated_at) values NULL," + s1 + "," + s2 + "," + s3 + "," + s4 + "," + s5 + "," + s6 + "," + s7 + ")")
	if err2 != nil {
		log.Fatalf("Can't add to Export Table %v", err2)
		return false, err2
	}

	return true, nil
}

func (c exporterDao) DelDE(id int64) (ok bool, goods_list []byte, goods_list_amount []byte, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, nil, nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return false, nil, nil, err1
	}

	//del good_list
	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select goods_list,goods_list_amount from export_table where id=" + s1)

	if err != nil {
		log.Fatalf("Can't search for Don Export", err)
		return false, nil, nil, err
	}

	for rows.Next() {
		err := rows.Scan(&goods_list, &goods_list_amount)
		if err != nil {
			log.Fatal("Can't scan for Don Export", err)
			return false, nil, nil, err
		}
		/*
			err1 := os.Remove(goods_list)
			if err1 != nil {
				log.Fatalf("Can't remove the old export good list", err)
				return false, err
			} */
	}

	//del in db
	s := "'" + strconv.FormatInt(id, 10) + "'"

	_, err2 := db.Exec("delete from export_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete in Export Table id number %d", id)
		return false, nil, nil, err2
	}

	return true, goods_list, goods_list_amount, nil
}

func (c exporterDao) UpdateDE(id int64, don *DonModel.DonExport) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(don.Id, 10) + "'"
	s1 := "'" + don.Created_by + "'"
	s2 := "'" + don.Reason + "'"
	s3 := "'" + string(don.Goods_list) + "'"
	s4 := "'" + string(don.Goods_list_amount) + "'"
	s5 := "'" + strconv.FormatInt(don.Total_money, 10) + "'"
	s6 := "'" + time.Now().String() + "'"

	_, err2 := db.Exec("update export_table set created_by=" + s1 + ", reason=" + s2 + ",goods_list=" + s3 + ",goods_list_amount=" + s4 + ", total_money=" + s5 + ",updated_at= " + s6 + "where id = " + s)
	if err2 != nil {
		log.Fatalf("Can't update export table %v", err2)
		return false, err2
	}
	return true, nil
}

func (c exporterDao) DEExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from export_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for Don Export", err)
		return false, err
	}

	id1 := 0
	for rows.Next() {
		err := rows.Scan(&id1)
		if err != nil {
			log.Fatal("Can't scan for Don Export", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c exporterDao) SelectAll() (r []*DonModel.DonExport, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from export_table")

	if err != nil {
		log.Fatalf("Can't search for all export", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Reason, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for export", err)
			return nil, err
		}
	}
	return r, nil
}

func (c exporterDao) SelectByID(id int64) (r *DonModel.DonExport, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from export_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for export by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Created_by, &r.Reason, &r.Goods_list, &r.Goods_list_amount, &r.Total_money, &r.Created_at, &r.Updated_at)
		if err != nil {
			log.Fatal("Can't scan for export by id", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

func (c exporterDao) SelectByCreatedAt(from_time int64, to_time int64) (r []*DonModel.DonExport, ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use exporter_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Export Database %v", err1)
		return nil, false, err
	}

	tin := time.Unix(from_time, 0)
	tin1 := tin.Format(time.UnixDate)

	s1 := "'" + tin1 + "'"

	tout := time.Unix(to_time, 0)
	tout1 := tout.Format(time.UnixDate)

	s2 := "'" + tout1 + "'"

	rows, err := db.Query("select * from export_table where (created_at <=" + s2 + ") and (" + s1 + "<= created_at)")

	if err != nil {
		log.Fatalf("Can't search for export by created_at", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Created_by, &r[i].Reason, &r[i].Goods_list, &r[i].Goods_list_amount, &r[i].Total_money, &r[i].Created_at, &r[i].Updated_at)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for sell", err)
			return nil, false, err
		}
	}
	return r, true, nil
}

var ExporterDao IExporter = exporterDao{}
