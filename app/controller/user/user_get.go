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

//func comment() {
//	var users []*pb.User
//	t := time.Now()
//	users = append(users, &pb.User{
//		Base: &pb.BaseUUID{
//			Id: uuid.New().String(),
//			Time: &pb.BaseDate{
//				CreatedAt: t.String()[:19],
//				UpdatedAt: t.String()[:19],
//				DeletedAt: t.String()[:19],
//			},
//		},
//		FullName:    "Firman Agam",
//		Gender:      "Male",
//		Email:       "agamwork28@gmail.com",
//		PhoneNumber: "081214025919",
//		Username:    "agambondan",
//		Password:    "agambondan",
//	})
//	findAllResponse := pb.FindAllResponse{
//		Response: &pb.Response{
//			Status:  true,
//			Message: "Data Found",
//			Error:   "",
//		},
//		Users: users,
//	}
//}
