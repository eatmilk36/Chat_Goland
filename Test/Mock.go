package Test

import (
	"chat/Common"
	"chat/Repositories/models"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type CryptoHelper struct {
	mock.Mock
}

func (c *CryptoHelper) Md5Hash(value string) string {
	args := c.Called(value)
	return args.String(0)
}

type RedisService struct {
	mock.Mock
}

func (r *RedisService) SaveUserLogin(ctx context.Context, username string, jwt string) error {
	args := r.Called(ctx, username, jwt)
	return args.Error(0)
}

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) CreateUser(user *models.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepository) GetUserByID(id uint) (*models.User, error) {
	args := u.Called(id)
	// args.Get(0) 是 *models.User，需要轉換類型
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepository) DeleteUser(id uint) error {
	args := u.Called(id)
	return args.Error(0)
}

func (u *UserRepository) GetUserByAccountAndPassword(account string, password string) (*models.User, error) {
	args := u.Called(account, password)
	user, ok := args.Get(0).(*models.User) // 取得模擬回傳的 *models.User
	if !ok {
		return nil, args.Error(1) // 如果模擬對象不存在，則回傳錯誤
	}
	return user, args.Error(1) // 回傳使用者與錯誤
}

type Jwt struct {
	mock.Mock
}

func (j *Jwt) GenerateJWT(username string) (string, error) {
	args := j.Called(username)
	return args.String(0), args.Error(1)
}

func (j *Jwt) ValidateJWT(tokenString string) (*Common.MyCustomClaims, error) {
	args := j.Called(tokenString)
	if claims := args.Get(0); claims != nil {
		return claims.(*Common.MyCustomClaims), args.Error(1)
	}
	return nil, args.Error(1)
}
