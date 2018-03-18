package dao

import (
	"database/sql"
	"log"
	DonModel "project1_update/others/models/reportmod"
	"strconv"
	"time"
)

type IHoursByDay interface {
	AddHour(hours *DonModel.HoursByDay) (bool, error)
	SelectAll() ([]*DonModel.HoursByDay, error)
	SelectByStaffID(id int64) ([]*DonModel.HoursByDay, bool, error)
}

type hoursByDayDao struct {
}

func (c hoursByDayDao) AddHour(hours *DonModel.HoursByDay) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use authen_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Authen  Database %v", err1)
		return false, err1
	}

	s1 := "'" + strconv.FormatInt(hours.StaffId, 10) + "'"

	tin := time.Unix(hours.In, 0)
	tin1 := tin.Format(time.UnixDate)

	s2 := "'" + tin1 + "'"

	tout := time.Unix(hours.Out, 0)
	tout1 := tout.Format(time.UnixDate)

	s3 := "'" + tout1 + "'"

	_, err2 := db.Exec("insert into hoursbyday_table (id,staffid,in,out) values NULL," + s1 + "," + s2 + "," + s3 + ")")
	if err2 != nil {
		log.Fatalf("Can't add to hoursbyday table %v", err2)
		return false, err2
	}

	return true, nil
}

func (c hoursByDayDao) SelectAll() (r []*DonModel.HoursByDay, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use authen_db")
	if err != nil {
		log.Fatalf("Can't connect to Authen Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from hoursbyday_table")

	if err != nil {
		log.Fatalf("Can't search for all hours", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].StaffId, &r[i].In, &r[i].Out)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for spend", err)
			return nil, err
		}
	}
	return r, nil
}

func (c hoursByDayDao) SelectByStaffID(id int64) (r []*DonModel.HoursByDay, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use authen_db")
	if err != nil {
		log.Fatalf("Can't connect to Authen Database %v", err1)
		return nil, false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from hoursbyday where staffid=" + s)

	if err != nil {
		log.Fatalf("Can't search for hoursbyday by id", err)
		return nil, false, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].StaffId, &r[i].In, &r[i].Out)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for hoursbyday", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

var HoursByDay IHoursByDay = hoursByDayDao{}
