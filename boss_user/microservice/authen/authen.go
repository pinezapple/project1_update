package authen

import (
	"log"

	Authen "project1_update/boss_user/microservice/authen/Authenrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AuthenAndClassify(staffid int64, passwd string, position int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10003", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to authen grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()
	c := Authen.NewAuthenClient(conn)

	r, err := c.AuthenAndClasify(context.Background(), &Authen.AuthenAndClasifyReq{Staffid: staffid, Passwd: passwd, Position: position})

	if err != nil {
		log.Fatalf("Can't authen from boss user %v", err)
		return false, err
	}

	return r.Ok, nil
}
