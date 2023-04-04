package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/retail-ai-test/internal/database"
	"github.com/retail-ai-test/internal/repo"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		if !isValidUserID(userID) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID"})
			c.Abort()
			return
		}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Authentication Failed",
			})
			c.Abort()
			return
		}
		isValid, err := isValidToken(c, token, userID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Authentication Failed",
			})
			c.Abort()
			return
		}

		if !isValid {
			c.JSON(http.StatusOK, gin.H{
				"message": "Authentication Failed",
			})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func isValidToken(c *gin.Context, token, uriUserID string) (bool, error) {
	token = strings.Replace(token, "Basic ", "", 1)
	rawDecodedText, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, err
	}
	st := strings.Split(string(rawDecodedText), ":")
	userID, password := st[0], st[1]
	if userID != uriUserID {
		return false, fmt.Errorf("userID is not matched")
	}
	userRepo := repo.NewUserRepo(db.GetConn())
	ok, err := userRepo.Validate(c, userID, password)
	if !ok {
		return false, err
	}
	return true, nil
}

func isValidUserID(input string) bool {
	minLength := 6
	maxLength := 20

	// Check if input meets criteria
	if len(input) < minLength || len(input) > maxLength {
		return false
	}
	for _, char := range input {
		if char == ' ' {
			return false
		}
	}
	return true
}
