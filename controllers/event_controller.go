package controllers

import (
	"net/http"
	"prakerja/eventmanagement/configs"
	"prakerja/eventmanagement/models"

	"github.com/labstack/echo/v4"
)

func InsertEventController(c echo.Context) error {
	var eventInput models.Event
	c.Bind(&eventInput)

	result := configs.DB.Create(&eventInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResp{
		Message: "success",
		Data:    eventInput,
	})
}

func GetEventController(c echo.Context) error {
	var events []models.Event
	result := configs.DB.Find(&events)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, models.BaseResp{
		Message: "Success",
		Data:    events,
	})
}

func DeleteEventController(c echo.Context) error {
	var events []models.Event
	id := c.Param("id")
	configs.DB.Delete(&events, id)
	return c.NoContent(http.StatusNoContent)
}
