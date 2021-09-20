package controllers

import (
	"crudgoeg/lib/database"
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

	return e.JSON(http.StatusOK, users)
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
	return e.JSON(http.StatusOK, result) 
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

	return e.JSON(http.StatusOK, users)
}

func UpdateUserController(e echo.Context) error {
	var userReq RequestUser
	e.Bind(&userReq)
	
	id, _ := strconv.Atoi(e.Param("id"))

	users, err := database.Update(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return e.JSON(http.StatusOK, users) 
}

func DeleteUserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))

	_ , err := database.Delete(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}