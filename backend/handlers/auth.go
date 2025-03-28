package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"eventhub/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginHandler(c echo.Context) error {
	var creds models.LoginCredentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный запрос"})
	}

	valid, err := database.ValidateUser(creds.Username, creds.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Неверный логин или пароль"})
		}
		c.Logger().Errorf("Ошибка при проверке данных пользователя: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при проверке данных пользователя"})
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Неверный логин или пароль"})
	}

	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		c.Logger().Errorf("Ошибка генерации токена: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка генерации токена"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func RegistrationHandler(c echo.Context) error {
	var creds models.RegistrationCredentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный запрос"})
	}

	if creds.Username == "" || creds.Password == "" || creds.FirstName == "" || creds.LastName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Заполнены не все поля"})
	}

	isUsernameTaken, err := database.IsUsernameTaken(creds.Username)
	if err != nil {
		c.Logger().Errorf("Ошибка при проверке существования пользователя: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при проверке существования пользователя"})
	}
	if isUsernameTaken {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Это имя пользователя уже занято"})
	}

	err = database.RegisterUser(creds.FirstName, creds.LastName, creds.Username, creds.Password)
	if err != nil {
		c.Logger().Errorf("Ошибка при создании пользователя: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при создании пользователя"})
	}

	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		c.Logger().Errorf("Ошибка при генерации токена: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при генерации токена"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Пользователь создан", "token": token})
}
