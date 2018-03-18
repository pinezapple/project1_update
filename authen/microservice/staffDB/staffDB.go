package staffDB

import (
	"log"
	"project1_update/authen/lib/MsgPack"
	Staff "project1_update/authen/microservice/services/StaffDBrpc"
	StaffModel "project1_update/others/models/staffmod"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AuthenAndClassify(staffid int64, passwd string, position int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("authen conn to staffDB err %v", err)
		return false, err
	}
	defer conn.Close()
	c := Staff.NewStaffClient(conn)

	r, err := c.SelectByStaffID(context.Background(), &Staff.SelectByIDReq{Id: staffid})

	if err != nil {
		log.Fatalf("Can't select staff by id in authen %v", err)
		return false, err
	}

	var staff *StaffModel.StaffMod
	err = MsgPack.Unmarshal(r.Payload, &staff)
	if err != nil {
		log.Fatalf("can't unmarshall in authen microservice %v", err)
		return false, err
	}
	if (staff.Id == staffid) && (staff.Passwd == passwd) && (staff.Position == position) {
		return true, nil
	}
	return false, nil
}
