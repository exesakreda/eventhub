package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"eventhub/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @Summary      Вход пользователя
// @Description  Аутентификация по логину и паролю, возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.LoginCredentials  true  "Логин и пароль"
// @Success      200          {object}  map[string]string         "JWT токен"
// @Failure      400          {object}  map[string]string         "Некорректный запрос"
// @Failure      401          {object}  map[string]string         "Неверный логин или пароль"
// @Failure      500          {object}  map[string]string         "Ошибка сервера"
// @Router       /login [post]
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

// @Summary      Регистрация пользователя
// @Description  Создание нового пользователя и возврат JWT токена
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.RegistrationCredentials  true  "Данные для регистрации"
// @Success      200          {object}  map[string]string               "Успешная регистрация и токен"
// @Failure      400          {object}  map[string]string               "Некорректный запрос или не все поля заполнены"
// @Failure      409          {object}  map[string]string               "Имя пользователя уже занято"
// @Failure      500          {object}  map[string]string               "Ошибка сервера"
// @Router       /register [post]
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
