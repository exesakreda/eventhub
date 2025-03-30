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

	// Id мероприятий, в которых участвует пользователь
	var joinedEventIds []int
	if err := database.DB.Table("event_participants").Where("user_id = ?", userId).Pluck("event_id", &joinedEventIds).Error; err != nil {
		c.Logger().Errorf("Ошибка при получении ID мероприятий: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	// Мероприятия, в которых участвует пользователь
	var joinedEvents []models.Event
	if err := database.DB.Where("id IN (?)", joinedEventIds).Find(&joinedEvents).Error; err != nil {
		c.Logger().Errorf("Ошибка при загрузке мероприятий пользователя: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	// Все открытые (доступные) мероприятия
	var openEvents []models.Event
	if err := database.DB.Where("is_public = ?", true).Find(&openEvents).Error; err != nil {
		c.Logger().Errorf("Ошибка при загрузке открытых мероприятий: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	return c.JSON(http.StatusOK, map[string][]models.Event{
		"joined_events": joinedEvents,
		"open_events":   openEvents,
	})
}

func JoinEvent(c echo.Context) error {
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

	var request struct {
		EventId uint `json:"event_id"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный JSON"})
	}

	var event models.Event
	if err := database.DB.First(&event, request.EventId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Мероприятие не найдено"})
		}
		c.Logger().Errorf("Ошибка при поиске мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	var existing models.EventParticipants
	if err := database.DB.Where("user_id = ? AND event_id = ?", userId, request.EventId).First(&existing).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Вы уже записаны на это мероприятие"})
	}

	eventParticipant := models.EventParticipants{UserID: userId, EventID: request.EventId}
	if err := database.DB.Create(&eventParticipant).Error; err != nil {
		c.Logger().Errorf("Ошибка при записи на мероприятие: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Пользователь записан на мероприятие"})
}

func QuitEvent(c echo.Context) error {
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

	var request struct {
		EventId uint `json:"event_id"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный JSON"})
	}

	var event models.Event
	if err := database.DB.First(&event, request.EventId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Мероприятие не найдено"})
		}
		c.Logger().Errorf("Ошибка при поиске мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	var existing models.EventParticipants
	if err := database.DB.Where("user_id = ? AND event_id = ?", userId, request.EventId).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Вы не записаны на это мероприятие"})
		}
		c.Logger().Errorf("Ошибка при отмене записи на мероприятие: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	if err := database.DB.Where("user_id = ? AND event_id = ?", userId, request.EventId).Delete(&models.EventParticipants{}).Error; err != nil {
		c.Logger().Errorf("Ошибка при удалении записи: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при отмене записи на мероприятие"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Запись на мероприятие отменена"})
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
