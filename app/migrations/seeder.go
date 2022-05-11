package migrations

import "github.com/agambondan/web-go-blog-grpc-rest/app/model"

var (
	roles model.Roles
)

// DataSeeds data to seeds
var DataSeeds []interface{} = []interface{}{
	roles.Seed(),
}
