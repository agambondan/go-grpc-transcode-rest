package model

type FirebaseToken struct {
	BaseInt
	Token                *string                `json:"token,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" binding:"required"`
	FirebaseTopics       *[]FirebaseTopic       `json:"firebase_topics,omitempty" gorm:"many2many:firebase_topic_token"`
	MessageNotifications *[]MessageNotification `json:"message_notifications,omitempty" gorm:"many2many:message_notification_token"`
}

type FirebaseTopic struct {
	BaseInt
	Topic                *string                `json:"topic,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" binding:"required"`
	FirebaseTokens       *[]FirebaseToken       `json:"firebase_tokens,omitempty" gorm:"many2many:firebase_topic_token"`
	MessageNotifications *[]MessageNotification `json:"message_notifications,omitempty" gorm:"many2many:message_notification_topic"`
}

type MessageNotification struct {
	BaseInt
	MessageNotificationBodyPointer
	FirebaseTokens *[]FirebaseToken `json:"firebase_tokens,omitempty" gorm:"many2many:message_notification_token"`
	FirebaseTopics *[]FirebaseTopic `json:"firebase_topics,omitempty" gorm:"many2many:message_notification_topic"`
}

type MessageNotificationBodyPointer struct {
	NotificationTitle *string `json:"notification_title,omitempty" gorm:"type:text;not null;" binding:"required"`
	NotificationBody  *string `json:"notification_body,omitempty" gorm:"type:text;not null;" binding:"required"`
	DataTitle         *string `json:"data_title," gorm:"type:text"`
	DataBody          *string `json:"data_body," gorm:"type:text"`
	ImageURL          *string `json:"image_url," gorm:"type:varchar(256)"`
}

type MessageNotificationAPI struct {
	Tokens *[]string `json:"tokens,omitempty"`
	Topics *[]string `json:"topics,omitempty"`
}

type MessageNotificationBody struct {
	NotificationTitle string `json:"notification_title,omitempty" gorm:"type:text;not null;"`
	NotificationBody  string `json:"notification_body,omitempty" gorm:"type:text;not null;"`
	DataTitle         string `json:"data_title," gorm:"type:text"`
	DataBody          string `json:"data_body" gorm:"type:text"`
	ImageURL          string `json:"image_url" gorm:"type:varchar(256)"`
}

type FCMRequest struct {
	MessageNotificationBody
	Data   map[string]string `json:"data,omitempty"`
	Topic  string            `json:"topic,omitempty"`
	Token  string            `json:"token,omitempty"`
	Tokens []string          `json:"tokens,omitempty"`
	Topics []string          `json:"topics,omitempty"`
}
