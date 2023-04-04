package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/model"
	"github.com/retail-ai-test/internal/model/request"
	"github.com/retail-ai-test/internal/model/response"
)

func (h *Handler) createUser(c *gin.Context) {
	var param request.CreateUserParam
	if err := c.ShouldBindJSON(&param); err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Account creation failed!",
			Cause:   "required user_id and password",
		})
		return
	}

	if param.NickName == "" {
		param.NickName = param.UserID
	}
	user := model.User{
		UserID:   param.UserID,
		Password: param.Password,
		NickName: param.NickName,
		Comment:  param.Comment,
	}

	ctx := c.Request.Context()
	result, err := h.Services.UserService.Create(ctx, user)
	if err != nil {
		internalErrorRes(c, response.ErrorMessage{
			Message: "Account creation failed",
			Cause:   "already same user_id is used",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account successfully created",
		"user":    result,
	})
}

func (h *Handler) getUserByID(c *gin.Context) {

	ctx := c.Request.Context()
	userID, _ := c.Get("userID")
	result, err := h.Services.UserService.FindByID(ctx, userID.(string))
	if err != nil {
		notFoundErrorRes(c, response.ErrorMessage{
			Message: "User details not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User details by user_id",
		"user":    result,
	})
}

func (h *Handler) updateUserByID(c *gin.Context) {
	userID, _ := c.Get("userID")

	var param request.UpdateUserParam
	if err := c.ShouldBindJSON(&param); err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "User update failed",
			Cause:   "required nickname or comment",
		})
		return
	}
	user := model.User{
		UserID:   userID.(string),
		NickName: param.NickName,
		Comment:  param.Comment,
	}

	ctx := c.Request.Context()
	result, err := h.Services.UserService.UpdateByID(ctx, user)
	if err != nil {
		internalErrorRes(c, response.ErrorMessage{
			Message: "User update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully updated",
		"user":    []*response.User{result},
	})
}

func (h *Handler) deleteUserByID(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _ := c.Get("userID")
	err := h.Services.UserService.DeleteByID(ctx, userID.(string))
	if err != nil {
		badRequestErrorRes(c, response.ErrorMessage{
			Message: "Account and user deletion failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Account and user successfully deleted",
	})
}
