package controllers

import (
	"crudgoeg/lib/database"
	"crudgoeg/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(e echo.Context) error {

	// id, _ := strconv.Atoi(e.QueryParam("id"))

	users, err := database.GetUsers()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return e.JSON(http.StatusOK, newResponseArray(*users))
}

func PostUserController(e echo.Context) error {
	var userReq RequestUser
	e.Bind(&userReq)

	result, err := database.StoreUser(userReq.toModel())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return e.JSON(http.StatusOK, newResponse(*result))
}

func GetbyIdUserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	users, err := database.GetbyIDUser(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return e.JSON(http.StatusOK, newResponse(*users))
}

func UpdateUserController(e echo.Context) error {
	var userReq RequestUser
	e.Bind(&userReq)

	id, _ := strconv.Atoi(e.Param("id"))

	result, err := database.Update(id, userReq.toModel())

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return e.JSON(http.StatusOK, newResponse(*result))
}

func DeleteUserController(e echo.Context) error {
	var user models.User

	id, _ := strconv.Atoi(e.Param("id"))
	_, errF := database.GetbyIDUser(id)
	_, err := database.Delete(id, user)

	if errF != nil {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}
