# RPC

This project is built for learning rpc basics

## Requirements:


>golang\
protoc\
protobuf\
grpc

### Instructions to run :( cloned repo or downloaded zip )

1. use latest version of go
2. go mod tidy
3. go run server.go
4. go run client/client.go ( you can write your client using .proto file or you can use the client provided ) 

### Instructions to run server :( .tar provided )

#### requirements
>docker

1. Download the .tar file throught the link.
2. docker load -i <path to the .tar file\>
3. docker run --rm -it -p 8000:8000 grpc-go
4. Now server got started

>Note: We have 10 users with id 1 to 10

### Test cases:
**input format: comma separated values ( id of the user ) (only for testing actual implementation may depend  on client application )**

**ouput : result or error**

---
**testcase 1 :**

input :\
4,10,2

output : \
4,10,2 \
Calling GetMultipleUserInfo procedure \
Received response => [Fname:"Kalyan"  City:"MA"  Phone:432567890  Height:5.7 Fname:"Ram"  City:"IN"  Phone:1823697890  Height:6.1 Fname:"Mary"  City:"LA"  Phone:1545787890  Height:5.8  Married:true] 

---

**testcase 2 :**

input :\
3,1

output: \
3,1 \
Calling GetMultipleUserInfo procedure \
Received response => [Fname:"Barghav"  City:"IN"  Phone:1875697890  Height:5.8  Married:true Fname:"Steve"  City:"LA"  Phone:1234567890  Height:5.8  Married:true] 

---

**testcase 3 :**

input :\
8

output : \
8 \
Calling GetSingleUserInfo procedure \
Received response => Fname:"Mariya"  City:"CN"  Phone:154532780  Height:5.4 


---

**testcase 4 :**

input :\
24

output : \
24\
Calling GetSingleUserInfo procedure\
found an error :  rpc error: code = Unknown desc = No data found related to the user

---

**testcase 5 :**

input :\
3,1,20,12

output : \
3,1,20,12\
Calling GetMultipleUserInfo procedure\
rpc error: code = InvalidArgument desc = users with the id [20 12] are not found
