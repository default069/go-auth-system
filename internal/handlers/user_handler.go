package handlers

import (
	"auto-system/internal/dto"
	"auto-system/internal/models"
	"auto-system/pkg/helpers"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type UserHandler struct{
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{}
}


func (h *UserHandler) SingUp(c *gin.Context){
	var req dto.SignUpRequest	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	hash, err := helpers.HashPassword(req.Password)
	if err != nil{
		helpers.InternalError(c)
		return
	}
	user := models.User{Name: req.Name, Email: req.Email,Password: hash}
	if err := h.DB.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey){
			helpers.Error(c ,http.StatusConflict, "email already exists")
			return
		}
		helpers.InternalError(c)
		return
	}
	c.JSON(http.StatusCreated, user)
}