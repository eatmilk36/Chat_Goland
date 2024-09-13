package Login

import (
	"chat/Common"
	"chat/Ineterface"
	"chat/Redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
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
	user, err := h.userRepo.GetUserByAccountAndPassword(req.Account, Common.Md5Hash(req.Password))

	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	jwt, _ := Common.GenerateJWT(user.Account)

	redisService := Redis.NewRedisService()
	err = redisService.SaveUserLogin(context.Background(), user.Account, jwt)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	c.JSON(http.StatusOK, jwt)
}
