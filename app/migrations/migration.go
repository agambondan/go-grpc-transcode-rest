package migrations

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/model"
)

// ModelMigrations models to migrate
var ModelMigrations []interface{} = []interface{}{
	&model.Role{},
	&model.User{},
	&model.Category{},
	&model.Article{},
	&model.FirebaseTopic{},
	&model.FirebaseToken{},
	&model.MessageNotification{},
	&model.Like{},
}
