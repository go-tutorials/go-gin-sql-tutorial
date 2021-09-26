package app

import "time"

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id,omitempty" firestore:"-"`
	Username    string     `json:"username,omitempty" gorm:"column:username" bson:"username,omitempty" dynamodbav:"username,omitempty" firestore:"username,omitempty" validate:"username,max=100"`
	Email       string     `json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty" validate:"required,email,max=100"`
	Url         string     `json:"url,omitempty" gorm:"column:url" bson:"url,omitempty" dynamodbav:"url,omitempty" firestore:"url,omitempty" validate:"required,url,max=255"`
	Phone       string     `json:"phone,omitempty" gorm:"column:phone" bson:"phone,omitempty" dynamodbav:"phone,omitempty" firestore:"phone,omitempty" validate:"required,phone,max=18"`
	Active      bool       `json:"active,omitempty" gorm:"column:active" bson:"active,omitempty" dynamodbav:"active,omitempty" firestore:"active,omitempty" true:"A" false:"D"`
	Locked      bool       `json:"locked,omitempty" gorm:"column:locked" bson:"locked,omitempty" dynamodbav:"locked,omitempty" firestore:"locked,omitempty" true:"1" false:"0"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty" gorm:"column:date_of_birth" bson:"dateOfBirth,omitempty" dynamodbav:"dateOfBirth,omitempty" firestore:"dateOfBirth,omitempty"`
}
