package user

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *Controller) FindAll(ctx context.Context, paginateRequest *pb.PaginateRequest) (*structpb.Value, error) {
	// init response message
	message := make(map[string]interface{})
	// get limit & offset
	limit := paginateRequest.GetSize()
	offset := paginateRequest.GetSize() * paginateRequest.GetPage()
	// find all user by limit & offset
	findAll, err := c.userRepository.FindAll(int(limit), int(offset))
	if err != nil {
		message["status"] = true
		message["message"] = ""
		message["error"] = err.Error()
		return structpb.NewValue(message)
	}
	// final response
	var responseUsers []interface{}
	lib.Merge(findAll, &responseUsers)
	message["status"] = true
	message["message"] = "Data found"
	message["error"] = ""
	message["data"] = responseUsers
	newMessage, err := structpb.NewValue(message)
	if err != nil {
		return nil, err
	}
	return newMessage, nil
}

//func (c *Controller) FindAll(ctx context.Context, paginateRequest *pb.PaginateRequest) (*pb.Response, error) {
//	limit := paginateRequest.GetSize()
//	offset := paginateRequest.GetSize() * paginateRequest.GetPage()
//	findAll, err := c.userRepository.FindAll(int(limit), int(offset))
//	users := &pb.Users{}
//	marshal, err := json.Marshal(findAll)
//	if err != nil {
//		log.Println(err)
//	}
//	lib.Merge(marshal, &users.Data)
//	responseUsers := &pb.Response_Users{
//		Users: users,
//	}
//	findAllResponse := pb.Response{
//		Data: responseUsers,
//	}
//	if err != nil {
//		findAllResponse.Error = err.Error()
//		findAllResponse.Status = false
//		findAllResponse.Message = fmt.Sprint("Data not found")
//	}
//	findAllResponse.Error = ""
//	findAllResponse.Status = true
//	findAllResponse.Message = fmt.Sprint("Data found")
//	lib.Merge(users, &findAllResponse.Data)
//	m, err := structpb.NewValue(map[string]interface{}{
//		"firstName": "John",
//		"lastName":  "Smith",
//		"isAlive":   true,
//		"age":       27,
//		"address": map[string]interface{}{
//			"streetAddress": "21 2nd Street",
//			"city":          "New York",
//			"state":         "NY",
//			"postalCode":    "10021-3100",
//		},
//		"phoneNumbers": []interface{}{
//			map[string]interface{}{
//				"type":   "home",
//				"number": "212 555-1234",
//			},
//			map[string]interface{}{
//				"type":   "office",
//				"number": "646 555-4567",
//			},
//		},
//		"children": []interface{}{},
//		"spouse":   nil,
//	})
//	if err != nil {
//		log.Println(err)
//	}
//	return &findAllResponse, nil
//}

//func (c *Controller) FindAll(ctx context.Context, paginateRequest *pb.PaginateRequest) (*pb.FindAllResponse, error) {
//	findAllResponse := pb.FindAllResponse{
//		Response: &pb.Response{},
//		Users:    []*pb.User{},
//	}
//	limit := paginateRequest.GetSize()
//	offset := paginateRequest.GetSize() * paginateRequest.GetPage()
//	findAll, err := c.userRepository.FindAll(int(limit), int(offset))
//	if err != nil {
//		findAllResponse.Response.Error = err.Error()
//		findAllResponse.Response.Status = false
//		findAllResponse.Response.Message = fmt.Sprint("Data not found")
//	}
//	findAllResponse.Response.Error = ""
//	findAllResponse.Response.Status = true
//	findAllResponse.Response.Message = fmt.Sprint("Data found")
//	lib.Merge(findAll, &findAllResponse.Users)
//	return &findAllResponse, nil
//}
