package user

import (
	"context"
	"fmt"
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
)

func (c *Controller) FindAll(ctx context.Context, paginateRequest *pb.PaginateRequest) (*pb.FindAllResponse, error) {
	findAllResponse := pb.FindAllResponse{
		Response: &pb.Response{},
		Users:    []*pb.User{},
	}
	limit := paginateRequest.GetSize()
	offset := paginateRequest.GetSize() * paginateRequest.GetPage()
	findAll, err := c.userRepository.FindAll(int(limit), int(offset))
	if err != nil {
		findAllResponse.Response.Error = err.Error()
		findAllResponse.Response.Status = false
		findAllResponse.Response.Message = fmt.Sprint("Data not found")
	}
	findAllResponse.Response.Error = ""
	findAllResponse.Response.Status = true
	findAllResponse.Response.Message = fmt.Sprint("Data found")
	lib.Merge(findAll, &findAllResponse.Users)
	return &findAllResponse, nil
}
