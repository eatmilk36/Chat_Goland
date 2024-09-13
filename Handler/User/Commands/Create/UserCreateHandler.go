package Create

import (
	"chat/Common"
	"chat/Ineterface"
	"chat/Repositories/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateHandler struct {
	userRepo Ineterface.UserRepository
}

// NewLoginHandler 建立 CreateHandler 並注入 UserRepository
func NewLoginHandler(userRepo Ineterface.UserRepository) *CreateHandler {
	return &CreateHandler{userRepo: userRepo}
}

func (h *CreateHandler) CreatUserCommand(c *gin.Context) {
	var req UserCreateRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.userRepo.CreateUser(&models.User{
		Account:     req.Account,
		Password:    Common.Md5Hash(req.Password),
		CreatedTime: req.Createdtime,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, "Create User Failed")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
