package handlers

import (
	"errors"
	"eventhub/database"
	"eventhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUserData(c echo.Context) error {
	username := c.Get("username").(string)

	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Пользователь не найден"})
		}
		c.Logger().Errorf("Ошибка при получении данных о пользователе: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при получении данных о пользователе"})
	}

	return c.JSON(http.StatusOK, user)
}
