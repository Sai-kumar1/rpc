package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/Sai-kumar1/totalitycorp/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserInfo struct {
	Id      int
	Fname   string
	City    string
	Phone   int32
	Height  float32
	Married bool
}

var userData map[int32]UserInfo = map[int32]UserInfo{
	1: {
		Id:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	},
	2: {
		Id:      2,
		Fname:   "Mary",
		City:    "LA",
		Phone:   1545787890,
		Height:  5.8,
		Married: true,
	},
	3: {
		Id:      3,
		Fname:   "Barghav",
		City:    "IN",
		Phone:   1875697890,
		Height:  5.8,
		Married: true,
	},
	4: {
		Id:      4,
		Fname:   "Kalyan",
		City:    "MA",
		Phone:   432567890,
		Height:  5.7,
		Married: false,
	},
	5: {
		Id:      5,
		Fname:   "Randi",
		City:    "IND",
		Phone:   457287890,
		Height:  5.0,
		Married: false,
	},
	6: {
		Id:      6,
		Fname:   "Ram",
		City:    "IN",
		Phone:   1852497890,
		Height:  5.2,
		Married: false,
	},
	7: {
		Id:      7,
		Fname:   "Carlos",
		City:    "LA",
		Phone:   126217890,
		Height:  5.5,
		Married: true,
	},
	8: {
		Id:      8,
		Fname:   "Mariya",
		City:    "CN",
		Phone:   154532780,
		Height:  5.4,
		Married: false,
	},
	9: {
		Id:      9,
		Fname:   "Akshaya",
		City:    "IN",
		Phone:   898197890,
		Height:  5.1,
		Married: true,
	},
	10: {
		Id:      10,
		Fname:   "Ram",
		City:    "IN",
		Phone:   1823697890,
		Height:  6.1,
		Married: false,
	},
}

type server struct{}

func (s *server) GetSingleUserInfo(ctx context.Context, request *userpb.SingleUserInfoRequest) (*userpb.SingleUserInfoResponse, error) {

	var id int32 = request.Id

	if _, doUserExist := userData[id]; !doUserExist {
		log.Println("GetSingleUserInfo:No user with the id ", id)
		return nil, errors.New("No data found related to the user")
	}
	response := &userpb.SingleUserInfoResponse{
		UserDetails: &userpb.UserInfo{
			Fname:   userData[id].Fname,
			City:    userData[id].City,
			Phone:   userData[id].Phone,
			Height:  userData[id].Height,
			Married: userData[id].Married,
		},
	}
	return response, nil
}

func (s *server) GetMultipleUserInfo(ctx context.Context, request *userpb.MultiUserInfoRequest) (*userpb.MultiUserInfoResponse, error) {
	var id []int32 = request.Id
	var data []*userpb.UserInfo
	var unfoundUsers []int32
	for _, userId := range id {
		if _, doUserExist := userData[userId]; !doUserExist {
			unfoundUsers = append(unfoundUsers, userId)
		} else {
			data = append(data, &userpb.UserInfo{
				Fname:   userData[userId].Fname,
				City:    userData[userId].City,
				Phone:   userData[userId].Phone,
				Height:  userData[userId].Height,
				Married: userData[userId].Married,
			})
		}
	}

	response := &userpb.MultiUserInfoResponse{
		UserDetails: data,
	}
	if len(unfoundUsers) != 0 {
		log.Println("GetMultipleUserInfo:found invalid user ids'",unfoundUsers)
		return response, status.Errorf(codes.InvalidArgument, "users with the id %v are not found", unfoundUsers)
	}
	return response, nil
}

func main() {

	lis, er := net.Listen("tcp", ":8000")
	if er != nil {
		log.Fatalf("error %v", er)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterGetUserInfoServiceServer(grpcServer, &server{})
	fmt.Println("grpc server is listening at localhost:8000")
	grpcServer.Serve(lis)
}
