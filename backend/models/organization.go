package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name       string `json:"name"`
	Founder_id int    `json:"founder_id"`
}

type OrganizationMember struct {
	UserId         int `gorm:"primaryKey" json:"user_id"`
	OrganizationId int `gorm:"primaryKey" json:"organization_id"`
}
