package Controller

import (
	"chat/Handler/Login"
	"chat/Repositories"
	"chat/Repositories/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// GetUser Login godoc
// @Summary User Login
// @Description Logs in a user with account and password credentials
// @Tags Login
// @Accept  json
// @Produce  json
// @Param LoginRequest body Login.LoginRequest true "Login credentials"
// @Success 200 {object} models.User "Successfully logged in user"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Router /login [post]
func (ctrl UserController) GetUser(c *gin.Context) {
	database := Repositories.GormUserRepository{}.InitDatabase()

	// 初始化 UserRepository
	userRepo := models.NewGormUserRepository(database)

	// 注入到 LoginHandler
	loginHandler := Login.NewLoginHandler(userRepo)

	// 呼叫 業務邏輯
	loginHandler.LoginQueryHandler(c)
}
