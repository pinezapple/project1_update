package staffmana

import (
	"log"
	staff "project1_update/boss_user/microservice/staffmana/Staffmanarpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func AddStaff(name string, passwd string, position int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("boss user conn to staffmana err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffmanaClient(conn)

	_, err1 := c.AddStaff(context.Background(), &staff.AddStaffReq{Name: name, Passwd: passwd, Position: position})
	if err1 != nil {
		log.Fatalf("can't add staff by boss user %v", err1)
		return false, err1
	}

	return true, err
}

func DelStaff(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("boss user conn to staffmana err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffmanaClient(conn)

	_, err1 := c.DelStaff(context.Background(), &staff.DelStaffReq{Id: id})
	if err1 != nil {
		log.Fatalf("can't del staff from boss user %v", err1)
		return false, err1
	}

	return true, err
}

func UpdateStaff(id int64, name string, passwd string, position int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("boss user conn to staffmana err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffmanaClient(conn)
	_, err1 := c.UpdateStaff(context.Background(), &staff.UpdateStaffReq{Id: id, Name: name, Passwd: passwd, Position: position})
	if err1 != nil {
		log.Fatalf("can't update staffDB from boss user %v", err1)
		return false, err1
	}

	return true, err
}

func SelectAllStaff() (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("boss user conn to staffmana err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := staff.NewStaffmanaClient(conn)

	r, err1 := c.SelectAllStaff(context.Background(), &staff.SelectAllStaffReq{})
	if err1 != nil {
		log.Fatalf("can't select all staff from boss user %v", err1)
		return nil, err1
	}

	return r.Payload, err
}

func SelectStaffById(id int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("boss user conn to staffmana err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := staff.NewStaffmanaClient(conn)

	r, err1 := c.SelectStaffById(context.Background(), &staff.SelectStaffByIdReq{Id: id})
	if err1 != nil {
		log.Fatalf("can't select staff by id from boss user %v", err1)
		return nil, err1
	}

	return r.Payload, err
}
