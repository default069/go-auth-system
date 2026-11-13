package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

func CheckPassword(hash, password string) bool{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil	
}

func Error (c *gin.Context, status int, msg string){
	c.AbortWithStatusJSON(status, gin.H{"error":msg})
}

func InternalError(c *gin.Context) {
	Error(c, http.StatusInternalServerError,"internal error")
}