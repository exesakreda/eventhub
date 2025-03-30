package handlers

import (
	"eventhub/database"
	"eventhub/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateOrganization(c echo.Context) error {
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

func JoinOrganization(c echo.Context) error {
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
