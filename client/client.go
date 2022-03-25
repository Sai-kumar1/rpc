package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Sai-kumar1/totalitycorp/userpb"
	"google.golang.org/grpc"
)

func main() {

	clientConn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	var input string
	

	client := userpb.NewGetUserInfoServiceClient(clientConn)
	for ;;{
		fmt.Println("enter the id with ',' separated \nex: 1,2,3 and quit to stop")
		fmt.Scanf("%s", &input)
		if input=="quit"{
			return
		}
		idFromSTDIN := strings.Split(input, ",")
		if len(idFromSTDIN) == 1 {
			v, er := strconv.Atoi(idFromSTDIN[0])
			if er != nil {
				log.Printf("found invalid input %v \n", v)
				return
			}
			fmt.Println("Calling GetSingleUserInfo procedure")
			request := &userpb.SingleUserInfoRequest{Id: int32(v)}
			resp, er := client.GetSingleUserInfo(context.Background(), request)
			if er != nil {
				fmt.Println("found an error : ", er)
			} else {
				fmt.Printf("Receive response => %v \n", resp.UserDetails)
	
			}
		} else if len(idFromSTDIN) > 1 {
			usersid := []int32{}
			for _, v := range idFromSTDIN {
				value, er := strconv.Atoi(v)
				if er != nil {
					log.Printf("found invalid input %v \n", v)
					return
				}
				usersid = append(usersid, int32(value))
			}
			// time.Sleep(2 * time.Second)
			fmt.Println("Calling GetMultipleUserInfo procedure")
	
			request2 := &userpb.MultiUserInfoRequest{Id: usersid}
			resp2, err2 := client.GetMultipleUserInfo(context.Background(), request2)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Printf("Receive response => %v \n", resp2.UserDetails)
			}
		} else {
			fmt.Println("enter valid number of inputs >0")
		}
	fmt.Println("--------------------------------------")

	}
	
}
