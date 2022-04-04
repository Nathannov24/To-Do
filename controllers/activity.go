package controllers

import (
	"net/http"
	"todo/lib/database"
	responses "todo/lib/response"
	"todo/models"

	echo "github.com/labstack/echo/v4"
)

func PostActivityController(c echo.Context) error {
	input := models.Activity{}
	c.Bind(&input)

	duplicate, _ := database.GetActivityByName(input.Name)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusDuplicated())
	}
	// Menyimpan data barang baru menggunakan fungsi InsertFeature
	_, e := database.PostActivity(input)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailedInput())
	}
	return c.JSON(http.StatusCreated, responses.StatusSuccessInput(input.Name))
}

func GetActivityController(c echo.Context) error {
	activity, err := database.GetAllActivity()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusInternalServerError())
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessGetData(activity))
}
