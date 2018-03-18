package staffDB

import (
	"log"
	staff "project1_update/staffmana/microservice/services/StaffDBrpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func AddStaff(name string, passwd string, position int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("staffmana conn to stafftomerDB err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffClient(conn)

	_, err1 := c.AddStaff(context.Background(), &staff.AddReq{Id: 0, Name: name, Passwd: passwd, Position: position})
	if err1 != nil {
		log.Fatalf("can't add staff to staffDB %v", err1)
		return false, err1
	}

	return true, err
}

func DelStaff(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("staffmana conn to staffDB err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffClient(conn)

	_, err1 := c.DelStaff(context.Background(), &staff.DelReq{Id: id})
	if err1 != nil {
		log.Fatalf("can't del staff from staffDB %v", err1)
		return false, err1
	}

	return true, err
}

func UpdateStaff(id int64, name string, passwd string, position int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("staffmana conn to staffDB err %v", err)
		return false, err
	}
	defer conn.Close()

	c := staff.NewStaffClient(conn)
	_, err1 := c.UpdateStaff(context.Background(), &staff.UpdateReq{Id: id, Name: name, Passwd: passwd, Position: position})
	if err1 != nil {
		log.Fatalf("can't update staffDB %v", err1)
		return false, err1
	}

	return true, err
}

func SelectAllStaff() (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("staffmana conn to staffDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := staff.NewStaffClient(conn)

	r, err1 := c.SelectAllStaff(context.Background(), &staff.SelectAllReq{})
	if err1 != nil {
		log.Fatalf("can't select all staff %v", err1)
		return nil, err1
	}

	return r.Payload, err
}

func SelectStaffById(id int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("staffmana conn to staffDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := staff.NewStaffClient(conn)

	r, err1 := c.SelectByStaffID(context.Background(), &staff.SelectByIDReq{Id: id})
	if err1 != nil {
		log.Fatalf("can't select staff by id %v", err1)
		return nil, err1
	}

	return r.Payload, err
}
