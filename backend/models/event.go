package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model  `swaggerignore:"true"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsPublic    bool   `json:"is_public"`
	Status      string `json:"status"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Location    string `json:"location"`
	CreatorId   uint   `json:"creator_id"`
}

type UpdateEventInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
	IsPublic    *bool   `json:"is_public"`
	Status      *string `json:"status"`
	Date        *string `json:"date"`
	StartTime   *string `json:"start_time"`
	EndTime     *string `json:"end_time"`
	Location    *string `json:"location"`
	CreatorId   *uint   `json:"creator_id"`
}

type EventParticipants struct {
	UserID  uint `gorm:"primaryKey" json:"user_id"`
	EventID uint `gorm:"primaryKey" json:"event_id"`
}
