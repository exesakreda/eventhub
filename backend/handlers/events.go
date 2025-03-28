package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateEventHandler(c echo.Context) error {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный запрос"})
	}

	err := database.DB.Create(&event).Error
	if err != nil {
		c.Logger().Errorf("Ошибка при создании мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при создании мероприятия"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Мероприятие создано"})
}

func GetEvents(c echo.Context) error {
	username := c.Get("username").(string)

	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Пользователь не найден"})
		}
		c.Logger().Errorf("Ошибка при запросе к базе данных: %v", result.Error)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}
	userId := user.ID
	status := c.QueryParam("status")
	role := c.QueryParam("role") // participant | creator
	orgId := c.QueryParam("organization_id")

	// if role != "participant" && role != "creator" {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректная роль"})
	// }
	// if orgId == "" {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Необходимо передать user_id или organization_id"})
	// }

	query := database.DB.Where("status = ?", status)

	if role == "participant" {
		query = query.Where("id IN (?)", database.DB.Table("event_participants").Select("event_id").Where("user_id = ?", userId))
	} else if role == "creator" {
		if orgId != "" {
			query = query.Where("creator_id =  ? AND organization_id = ?", userId, orgId)
		} else {
			query = query.Where("creator_id =  ?", userId)
		}
	}

	var events []models.Event
	err := query.Find(&events).Error
	if err != nil {
		c.Logger().Errorf("Ошибка при поиске мероприятий: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	if len(events) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Мероприятия не найдены"})
	}

	return c.JSON(http.StatusOK, events)
}

func UpdateEvent(c echo.Context) error {
	id := c.QueryParam("event_id")

	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Мероприятие не найдено"})
	}

	var input models.UpdateEventInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный запрос"})
	}

	if input.Title != nil {
		event.Title = *input.Title
	}
	if input.Description != nil {
		event.Description = *input.Description
	}
	if input.Category != nil {
		event.Category = *input.Category
	}
	if input.IsPublic != nil {
		event.IsPublic = *input.IsPublic
	}
	if input.Status != nil {
		event.Status = *input.Status
	}
	if input.Date != nil {
		event.Date = *input.Date
	}
	if input.StartTime != nil {
		event.StartTime = *input.StartTime
	}
	if input.EndTime != nil {
		event.EndTime = *input.EndTime
	}
	if input.Location != nil {
		event.Location = *input.Location
	}

	if err := database.DB.Save(&event).Error; err != nil {
		c.Logger().Errorf("Ошибка при обновлении мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при обновлении мероприятия"})
	}

	return c.NoContent(http.StatusOK)
}

func DeleteEvent(c echo.Context) error {
	eventId := c.QueryParam("event_id")

	err := database.DB.Delete(&models.Event{}, "id = ?", eventId).Error
	if err != nil {
		c.Logger().Errorf("Ошибка при удалении мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	return c.JSON(http.StatusOK, "Мероприятие удалено")
}
