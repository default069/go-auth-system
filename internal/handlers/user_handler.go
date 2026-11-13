package handlers

import (
	"auto-system/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpRequest struct{
	Email 	 string 	`json:"email" binding:"required,email"`
	Name  	 string		`json:"name"  binding:"required,min=2,max=100"`
	Password string		`json:"password" binding:"required,min=8"`
}

type UserHandler struct{
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{}
}


func (h *UserHandler) SingUp(c *gin.Context){
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	
	user := models.User{Name: req.Name, Email:req.Email,Password:string(hash)}
	if err := h.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal error"})
		return
	}
	c.JSON(http.StatusCreated, user)
}