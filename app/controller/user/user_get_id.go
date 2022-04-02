package user

import (
	"context"
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	"github.com/agambondan/web-go-blog-grpc-rest/app/model"
	pb "github.com/agambondan/web-go-blog-grpc-rest/grpc/gen/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *Controller) FindByID(ctx context.Context, paginateRequest *pb.PaginateRequest) (*structpb.Value, error) {
	baseResponse := model.BaseResponse{}
	// init response message
	message := make(map[string]interface{})
	// find all user by limit & offset
	uuidParse, err := uuid.Parse(paginateRequest.GetUuid())
	if err != nil {
		baseResponse.Failed(err.Error(), "Data not found", 404)
		return structpb.NewValue(baseResponse.ConvertToMap())
	}
	findByID, err := c.userRepository.FindById(lib.UUIDPtr(uuidParse))
	if err != nil {
		baseResponse.Failed(err.Error(), "Data not found", 404)
		return structpb.NewValue(baseResponse.ConvertToMap())
	} else {
		baseResponse.Success("Data found")
		message = baseResponse.ConvertToMap()
	}
	// final response
	var responseUser map[string]interface{}
	lib.Merge(findByID, &responseUser)
	message["data"] = responseUser
	return structpb.NewValue(message)
}
