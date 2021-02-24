package handler

import (
	"net/http"

	"github.com/xuanbo/pig/entity"
	"github.com/xuanbo/pig/model"

	"github.com/labstack/echo/v4"
)

// User 用户API
type User struct {
}

// Login 登录
func (u *User) Login(ctx echo.Context) error {
	var user entity.User
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, model.OK(""))
}
