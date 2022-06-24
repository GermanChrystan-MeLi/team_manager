package handler

import (
	"context"

	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/internal/user"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/web"
	"github.com/gin-gonic/gin"
)

//=================================================================================//
type User struct {
	userService user.UserService
}

//=================================================================================//
func NewUser(us user.UserService) *User {
	return &User{
		userService: us,
	}
}

//=================================================================================//
func (u *User) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		userID := c.Request.Header.Get("user_id")
		profile, err := u.userService.GetProfile(ctx, userID)
		if err != nil {
			web.Error(c, 400, err.Error())
			return
		} else {
			web.Success(c, 200, profile)
			return
		}
	}
}

//=================================================================================//
func (u *User) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var userRegister dto.UserRegister
		if bindingErr := c.ShouldBindJSON(&userRegister); bindingErr != nil {
			web.Error(c, 400, bindingErr.Error())
			return
		} else {
			session, err := u.userService.Register(ctx, userRegister)
			if err != nil {
				web.Error(c, 400, err.Error())
				return
			}
			web.Success(c, 201, session)
			return
		}
	}
}

//=================================================================================//
func (u *User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var userLogin dto.UserLogin
		if bindingErr := c.ShouldBindJSON(&userLogin); bindingErr != nil {
			web.Error(c, 400, bindingErr.Error())
			return
		} else {
			session, err := u.userService.Login(ctx, userLogin)
			if err != nil {
				web.Error(c, 400, err.Error())
				return
			}
			web.Success(c, 201, session)
			return
		}
	}
}

//=================================================================================//
func (u *User) EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var userEdit dto.UserEdit
		if bindingErr := c.ShouldBindJSON(&userEdit); bindingErr != nil {
			web.Error(c, 400, bindingErr.Error())
			return
		} else {
			session, err := u.userService.EditUser(ctx, userEdit)
			if err != nil {
				web.Error(c, 400, err.Error())
				return
			}
			web.Success(c, 200, session)
			return
		}
	}
}

//=================================================================================//
func (u *User) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		id := c.Param("id")
		err := u.userService.DeleteUser(ctx, id)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		} else {
			web.Success(c, 204, "user deleted")
			return
		}
	}
}
