package dao

import (
	"database/sql"
	"log"
	staffModel "project1_update/others/models/staffmod"
	"strconv"
)

type Istaff interface {
	AddStaff(staff *staffModel.StaffMod) (bool, error)
	DelStaff(id int64) (bool, error)
	UpdateStaff(id int64, new *staffModel.StaffMod) (bool, error)
	StaffExist(id int64) (bool, error)
	SelectAll() ([]*staffModel.StaffMod, error)
	SelectByID(id int64) (*staffModel.StaffMod, bool, error)
}

type staffDao struct {
}

func (c staffDao) AddStaff(staff *staffModel.StaffMod) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Staff Database %v", err1)
		return false, err1
	}
	s := "'" + staff.Name + "'"
	s1 := "'" + staff.Passwd + "'"
	s2 := "'" + strconv.FormatInt(staff.Position, 10) + "'"
	_, err2 := db.Exec("insert into staff_table (id,name,passwd,position) values NULL," + s + "," + s1 + "," + s2 + ")")
	if err2 != nil {
		log.Fatalf("Can't add staff %v", err2)
		return false, err2
	}

	return true, nil
}

func (c staffDao) DelStaff(id int64) (ok bool, err error) {

	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Staff Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	_, err2 := db.Exec("delete from staff_table where id =" + s)
	if err2 != nil {
		log.Fatalf("Can't delete staff id number %d", id)
		return false, err2
	}

	return true, nil
}

func (c staffDao) UpdateStaff(id int64, new *staffModel.StaffMod) (ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Staff Database %v", err1)
		return false, err1
	}

	s1 := "'" + strconv.FormatInt(id, 10) + "'"
	s2 := "'" + new.Name + "'"
	s3 := "'" + new.Passwd + "'"
	s4 := "'" + strconv.FormatInt(new.Position, 10) + "'"

	_, err2 := db.Exec("update staff_table set name=" + s2 + ", passwd=" + s3 + ", position=" + s4 + "where id = " + s1)
	if err2 != nil {
		log.Fatalf("Can't update staff_db %v", err2)
		return false, err2
	}
	return true, nil
}

func (c staffDao) StaffExist(id int64) (ok bool, er error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err1 != nil {
		log.Fatalf("Can't connect to Staff Database %v", err1)
		return false, err1
	}

	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select id from staff_table where id=" + s)

	if err != nil {
		log.Fatalf("Can't search for staff", err)
		return false, err
	}
	id1 := 0
	for rows.Next() {
		err := rows.Scan(id1)
		if err != nil {
			log.Fatal("Can't scan for staff", err)
			return false, err
		}
		if id1 == 0 {
			return false, err
		}
	}
	return true, err
}

func (c staffDao) SelectAll() (r []*staffModel.StaffMod, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err != nil {
		log.Fatalf("Can't connect to staff Database %v", err1)
		return nil, err1
	}

	rows, err := db.Query("select * from staff_table")

	if err != nil {
		log.Fatalf("Can't search for all staff", err)
		return nil, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&r[i].Id, &r[i].Name, &r[i].Passwd, &r[i].Position)
		i = i + 1
		if err != nil {
			log.Fatal("Can't scan for staff", err)
			return nil, err
		}
	}
	return r, nil
}

func (c staffDao) SelectByID(id int64) (r *staffModel.StaffMod, ok bool, err error) {
	db, err := sql.Open("mysql", "root:Tungtinhte@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Can't connect to mysql %v", err)
		return nil, false, err
	}
	defer db.Close()

	_, err1 := db.Exec("use staff_db")
	if err1 != nil {
		log.Fatalf("Can't connect to staff Database %v", err)
		return nil, false, err
	}
	s := "'" + strconv.FormatInt(id, 10) + "'"
	rows, err := db.Query("select * from staff_table where id =" + s)

	if err != nil {
		log.Fatalf("Can't search for staff by id %v", err)
		return nil, false, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Id, &r.Name, &r.Passwd, &r.Position)
		if err != nil {
			log.Fatalf("Can't scan for staff %v", err)
			return nil, false, err
		}
	}
	return r, false, nil
}

var StaffDao Istaff = staffDao{}
