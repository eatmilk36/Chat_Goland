package Login

import (
	"chat/Ineterface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	userRepo Ineterface.UserRepository
}

// NewLoginHandler 建立 LoginHandler 並注入 UserRepository
func NewLoginHandler(userRepo Ineterface.UserRepository) *LoginHandler {
	return &LoginHandler{userRepo: userRepo}
}

func (h *LoginHandler) LoginQueryHandler(c *gin.Context) {
	var req LoginRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 使用解析出的 account 和 password
	user, err := h.userRepo.GetUserByAccountAndPassword(req.Account, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, "user not found")
	}

	c.JSON(http.StatusOK, user)
}
