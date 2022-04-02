package model

type Role struct {
	BaseInt
	Name *string `json:"name,omitempty" gorm:"type:varchar(16);not null;index:idx_role_name_deleted_at,unique,where:deleted_at is null"`
}
