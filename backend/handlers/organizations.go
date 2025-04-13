package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateOrganizationHandler godoc
// @Summary Создание организации
// @Description Создает новую организацию
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization body models.Organization true "Информация об организации"
// Success 200 {object} map[string]string {"success": "Организация создана"}
// Failure 400 {object} map[string]string {"error": "Некорректный запрос"}
// Failure 500 {object} map[string]string {"error": "Ошибка при создании организации"}
// Router /organizations/create [post]
func CreateOrganizationHandler(c echo.Context) error {
	var organization models.Organization

	if err := c.Bind(&organization); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный запрос"})
	}

	err := database.DB.Create(&organization).Error
	if err != nil {
		c.Logger().Errorf("Ошибка при создании организации: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при создании организации"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Организация создана"})
}

// JoinOrganizationHandler godoc
// @Summary Присоединение пользователя к организации
// @Description Позволяет пользователю присоединиться к организации по заданным ID
// @Tags organizations
// @Accept json
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Param organization_id query int true "ID организации"
// Success 200 {object} map[string]string {"success": "Пользователь присоединился к организации"}
// Failure 400 {object} map[string]string {"error": "Некорректный user_id"}
// Failure 400 {object} map[string]string {"error": "Некорректный organization_id"}
// Failure 404 {object} map[string]string {"error": "Организация не найдена"}
// Failure 500 {object} map[string]string {"error": "Ошибка при присоединении к организации"}
// @Router /organizations/join [post]
func JoinOrganizationHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некореектный user_id"})
	}

	organizationId, err := strconv.Atoi(c.QueryParam("organization_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некореектный organization_id"})
	}

	var organization models.Organization
	if err := database.DB.First(&organization, organizationId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Организация не найдена"})
	}

	member := models.OrganizationMember{UserId: userId, OrganizationId: organizationId}
	if err := database.DB.Create(&member).Error; err != nil {
		c.Logger().Errorf("Ошибка при присоединении к организации: %v", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при присоединении к организации")
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Пользователь присоединился к организации"})
}

// QuitOrganizationHandler godoc
// @Summary Выход пользователя из организации
// @Description Позволяет пользователю выйти из организации. Если пользователь является основателем организации, она будет удалена.
// @Tags organizations
// @Accept json
// @Produce json
// @Param user_id query int true "ID пользователя"
// @Param organization_id query int true "ID организации"
// Success 200 {object} map[string]string {"success": "Пользователь вышел из организации"}
// Failure 400 {object} map[string]string {"error": "Некорректный user_id"}
// Failure 400 {object} map[string]string {"error": "Некорректный organization_id"}
// Failure 404 {object} map[string]string {"error": "Организация не найдена"}
// Failure 500 {object} map[string]string {"error": "Ошибка при выходе из организации"}
// Failure 500 {object} map[string]string {"error": "Ошибка при удалении организации"}
// @Router /organizations/quit [post]
func QuitOrganizationHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некореектный user_id"})
	}

	organizationId, err := strconv.Atoi(c.QueryParam("organization_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некореектный organization_id"})
	}

	var organization models.Organization
	if err := database.DB.First(&organization, organizationId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Организация не найдена"})
	}

	member := models.OrganizationMember{UserId: userId, OrganizationId: organizationId}

	if err := database.DB.Delete(&member).Error; err != nil {
		c.Logger().Errorf("Ошибка при выходе из организации: %v", err)
		return c.JSON(http.StatusInternalServerError, "Ошибка при выходе из организации")
	}

	if organization.Founder_id == userId {
		if err := database.DB.Delete(&organization).Error; err != nil {
			c.Logger().Errorf("Ошибка при удалении организации: %v", err)
			return c.JSON(http.StatusInternalServerError, "Ошибка при удалении организации")
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "Пользователь вышел из организации"})
}
