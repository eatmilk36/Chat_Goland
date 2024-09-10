package Login

import (
	"chat/Repositories/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var req LoginRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 使用解析出的 account 和 password
	user, err := models.GetUserByAccountAndPassword(req.Account, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, "user not found")
	}

	c.JSON(http.StatusOK, user)
}
