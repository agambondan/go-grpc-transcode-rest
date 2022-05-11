package model

import "github.com/agambondan/web-go-blog-grpc-rest/app/lib"

type Role struct {
	BaseInt
	Name *string `json:"name,omitempty" gorm:"type:varchar(16);not null;index:idx_role_name_deleted_at,unique,where:deleted_at is null"`
}

type Roles []Role

func (r *Roles) Seed() *Roles {
	name := []string{"admin", "writer", "reader"}
	for i := 0; i < len(name); i++ {
		var role Role
		role.Name = lib.Strptr(name[i])
		*r = append(*r, role)
	}
	return r
}
