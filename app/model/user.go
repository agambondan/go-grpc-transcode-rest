package model

type User struct {
	BaseUUID
	BaseImage
	UserAPI
}

type UserAPI struct {
	FullName    *string `json:"full_name,omitempty" gorm:"type:varchar(48);not null;"`
	FirstName   *string `json:"first_name,omitempty" gorm:"type:varchar(24);not null;"`
	LastName    *string `json:"last_name,omitempty" gorm:"type:varchar(24);not null;"`
	Gender      *string `json:"gender,omitempty" gorm:"type:varchar(24);not null;"`
	Email       *string `json:"email,omitempty" gorm:"type:varchar(64);not null;index:idx_email_deleted_at,unique,where:deleted_at is null"`
	PhoneNumber *string `json:"phone_number,omitempty" gorm:"type:varchar(14);not null;index:idx_phone_number_deleted_at,unique,where:deleted_at is null"`
	Username    *string `json:"username,omitempty" gorm:"type:varchar(36);not null;index:idx_username_deleted_at,unique,where:deleted_at is null"`
	Password    *string `json:"password,omitempty" gorm:"type:varchar(256);not null;"`
	Instagram   *string `json:"instagram,omitempty" gorm:"type:varchar(24)"`
	Facebook    *string `json:"facebook,omitempty" gorm:"type:varchar(24)"`
	Twitter     *string `json:"twitter,omitempty" gorm:"type:varchar(24)"`
	LinkedIn    *string `json:"linked_in,omitempty" gorm:"type:varchar(24)"`
	RoleId      *int64  `json:"role_id,omitempty" gorm:"type:smallint;not null;"`
}
