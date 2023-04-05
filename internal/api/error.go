package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/internal/pkg/model/response"
)

func forbiddenErrorRes(c *gin.Context, message response.ErrorMessage) {
	c.JSON(http.StatusForbidden, message)
}

func badRequestErrorRes(c *gin.Context, message response.ErrorMessage) {
	c.JSON(http.StatusBadRequest, message)
}

func internalErrorRes(c *gin.Context, message response.ErrorMessage) {
	c.JSON(http.StatusInternalServerError, message)
}

func notFoundErrorRes(c *gin.Context, message response.ErrorMessage) {
	c.JSON(http.StatusNotFound, message)
}
