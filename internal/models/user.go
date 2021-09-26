package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id           string        `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" validate:"required,max=40"`
	Username     string        `json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email        string        `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" validate:"email,max=100"`
	Phone        string        `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" validate:"required,phone,max=18"`
	DateOfBirth  *time.Time    `json:"dateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth"`
	Interests    []string      `mapstructure:"interests" json:"interests,omitempty" gorm:"column:interests" bson:"interests,omitempty" dynamodbav:"interests,omitempty" firestore:"interests,omitempty"`
	Skills       []SkillItem   `mapstructure:"skills" json:"skills,omitempty" gorm:"column:skills" bson:"skills,omitempty" dynamodbav:"skills,omitempty" firestore:"skills,omitempty"`
	Achievements []Achievement `mapstructure:"achievements" json:"achievements,omitempty" gorm:"column:achievements" bson:"achievements,omitempty" dynamodbav:"achievements,omitempty" firestore:"achievements,omitempty"`
	Settings     *UserSettings `mapstructure:"settings" json:"settings,omitempty" gorm:"type:settings;column:settings" bson:"settings,omitempty" dynamodbav:"settings,omitempty" firestore:"settings,omitempty"`
}

type SkillItem struct {
	Skill   string `mapstructure:"skill" json:"skill,omitempty" gorm:"column:skill" bson:"skill,omitempty" dynamodbav:"skill,omitempty" firestore:"skill,omitempty"`
	Hirable bool   `mapstructure:"hirable" json:"hirable,omitempty" gorm:"column:hirable" bson:"hirable,omitempty" dynamodbav:"hirable,omitempty" firestore:"hirable,omitempty"`
}

func (i SkillItem) Value() (driver.Value, error) {
	//bs, err := json.Marshal(i)
	return json.Marshal(i)
}

func (i SkillItem) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &i)
}

type Achievement struct {
	Subject     string `mapstructure:"subject" gorm:"column:subject" json:"subject,omitempty" gorm:"subject" bson:"subject,omitempty" dynamodbav:"subject,omitempty" firestore:"subject,omitempty"`
	Description string `mapstructure:"description" gorm:"column:description" json:"description,omitempty" gorm:"description" bson:"description,omitempty" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
}

type AchievementSlice []Achievement

func (a Achievement) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	return string(b), err
}

func (a *Achievement) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, a)
	case string:
		return json.Unmarshal([]byte(v), a)
	}
	return errors.New("type assertion failed")
	//b, ok := value.([]byte)
	//if !ok {
	//	return errors.New("type assertion to []byte failed")
	//}
	//return json.Unmarshal(b, &a)
}

type UserSettings struct {
	UserId         string `mapstructure:"id" json:"-" gorm:"column:id" bson:"-" dynamodbav:"-" firestore:"-"`
	Language       string `mapstructure:"language" json:"language,omitempty" gorm:"column:language" bson:"language,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty"`
	DateFormat     string `mapstructure:"date_format" json:"dateFormat,omitempty" gorm:"column:dateformat" bson:"dateFormat,omitempty" dynamodbav:"dateFormat,omitempty" firestore:"dateFormat,omitempty"`
	DateTimeFormat string `mapstructure:"date_time_format" json:"dateTimeFormat,omitempty" gorm:"column:datetimeformat,omitempty" bson:"dateTimeFormat,omitempty" dynamodbav:"dateTimeFormat,omitempty" firestore:"dateTimeFormat,omitempty"`
	TimeFormat     string `mapstructure:"time_format" json:"timeFormat,omitempty" gorm:"column:timeformat" bson:"timeFormat,omitempty" dynamodbav:"timeFormat,omitempty" firestore:"timeFormat,omitempty"`
	Notification   bool   `mapstructure:"notification" json:"notification,omitempty" gorm:"column:notification" bson:"notification,omitempty" dynamodbav:"notification,omitempty" firestore:"notification,omitempty"`
}

func (u UserSettings) Value() (driver.Value, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func (u *UserSettings) Scan(input interface{}) error {
	bytes, ok := input.([]byte)
	fmt.Println("This is in scan progress")
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", input))
	}
	return json.Unmarshal(bytes, &u)
}
