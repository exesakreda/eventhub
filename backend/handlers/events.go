package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventIDRequest struct {
	EventId uint `json:"event_id"`
}

type SuccessResponse struct {
	Success string `json:"success"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Создание мероприятия
// @Description Хендлер для создания нового мероприятия
// @Tags Events
// @Accept json
// @Produce json
// @Param event body models.Event true "Мероприятие"
// Success 200 {object} map[string]string{"success": "Мероприятие создано"}
// Failure 400 {object} map[string]string{"error": "Некорректный запрос"}
// Failure 500 {object} map[string]string{"error": "Ошибка при создании мероприятия"}
// @Router /event/create [post]
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

// @Summary Получение списка мероприятий
// @Description Хендлер для получения мероприятий, в которых участвует пользователь, а также всех открытых мероприятий
// @Tags Events
// @Produce json
// Success 200 {object} map[string][]models.Event{"joined_events": []models.Event, "open_events": []models.Event}
// Failure 404 {object} map[string]string{"error": "Пользователь не найден"}
// Failure 500 {object} map[string]string{"error": "Ошибка сервера"}
// @Router /user/getevents [get]
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

	var joinedEventIds []int
	if err := database.DB.Table("event_participants").Where("user_id = ?", userId).Pluck("event_id", &joinedEventIds).Error; err != nil {
		c.Logger().Errorf("Ошибка при получении ID мероприятий: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	var joinedEvents []models.Event
	if err := database.DB.Where("id IN (?)", joinedEventIds).Find(&joinedEvents).Error; err != nil {
		c.Logger().Errorf("Ошибка при загрузке мероприятий пользователя: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

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

// @Summary Записаться на мероприятие
// @Description Хендлер для записи пользователя на мероприятие
// @Tags Events
// @Accept json
// @Produce json
// @Param event_id body EventIDRequest true "ID мероприятия"
// Success 200 {object} map[string]string{"success": "Пользователь записан на мероприятие"}
// Failure 400 {object} map[string]string{"error": "Некорректный JSON"}
// Failure 404 {object} map[string]string{"error": "Мероприятие не найдено"}
// Failure 409 {object} map[string]string{"error": "Вы уже записаны на это мероприятие"}
// Failure 500 {object} map[string]string{"error": "Ошибка сервера"}
// @Router /event/join [post]
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

// @Summary Отказаться от участия в мероприятии
// @Description Хендлер для отмены записи пользователя на мероприятие
// @Tags Events
// @Accept json
// @Produce json
// @Param event_id body EventIDRequest true "ID мероприятия"
// Success 200 {object} map[string]string{"success": "Запись на мероприятие отменена"}
// Failure 400 {object} map[string]string{"error": "Некорректный JSON"}
// Failure 404 {object} map[string]string{"error": "Мероприятие не найдено"}
// Failure 409 {object} map[string]string{"error": "Вы не записаны на это мероприятие"}
// Failure 500 {object} map[string]string{"error": "Ошибка сервера"}
// @Router /event/quit [post]
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

// @Summary Обновить информацию о мероприятии
// @Description Хендлер для обновления данных мероприятия
// @Tags Events
// @Accept json
// @Produce json
// @Param event_id query int true "ID мероприятия"
// @Param event body models.UpdateEventInput true "Данные для обновления"
// Success 200 {object} map[string]string{"success": "Мероприятие обновлено"}
// Failure 400 {object} map[string]string{"error": "Некорректный запрос"}
// Failure 404 {object} map[string]string{"error": "Мероприятие не найдено"}
// Failure 500 {object} map[string]string{"error": "Ошибка при обновлении мероприятия"}
// @Router /event/update [put]
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

// @Summary Удалить мероприятие
// @Description Хендлер для удаления мероприятия по ID
// @Tags Events
// @Accept json
// @Produce json
// @Param event_id query int true "ID мероприятия"
// Success 200 {object} map[string]string{"success": "Мероприятие удалено"}
// Failure 404 {object} map[string]string{"error": "Мероприятие не найдено"}
// Failure 500 {object} map[string]string{"error": "Ошибка сервера"}
// @Router /event/delete [delete]
func DeleteEvent(c echo.Context) error {
	eventId := c.QueryParam("event_id")

	err := database.DB.Delete(&models.Event{}, "id = ?", eventId).Error
	if err != nil {
		c.Logger().Errorf("Ошибка при удалении мероприятия: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка сервера"})
	}

	return c.JSON(http.StatusOK, "Мероприятие удалено")
}
