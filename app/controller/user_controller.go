package controller

import (
	"net/http"
	"point-of-sale/app/middleware"
	"point-of-sale/app/model"
	"point-of-sale/config"
	"point-of-sale/utils/dto"

	"github.com/labstack/echo/v4"
)

func LoginCashier(c echo.Context) error {
	var (
		request dto.LoginRequest
		user    model.User
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := config.Db.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if user.Password != request.Password {
		return c.JSON(http.StatusInternalServerError, "incorrect password")
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)

}

func LoginAdmin(c echo.Context) error {
	var (
		request dto.LoginRequest
		user    model.User
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := config.Db.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if user.Password != request.Password {
		return c.JSON(http.StatusInternalServerError, "incorrect password")
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
