package service

import (
	"context"

	"github.com/dath-251-thuanle/file-sharing-web-backend2/config"
	"github.com/dath-251-thuanle/file-sharing-web-backend2/internal/domain"
	"github.com/dath-251-thuanle/file-sharing-web-backend2/pkg/utils"
	"github.com/gin-gonic/gin"
)

type TOTPSetupResponse struct {
	Secret string `json:"secret"`
	QRCode string `json:"qrCode"`
}

type UserService interface {
	GetUserById(id string) (*domain.User, *utils.ReturnStatus)
	GetUserByEmail(email string) (*domain.User, *utils.ReturnStatus)
}

type AuthService interface {
	CreateUser(username, password, email string) (*domain.User, *utils.ReturnStatus)
	Login(email, password string) (user *domain.User, accessToken string, err *utils.ReturnStatus)
	SetupTOTP(userID string) (*TOTPSetupResponse, *utils.ReturnStatus)
	VerifyTOTP(userID string, code string) (bool, *utils.ReturnStatus)
	Logout(ctx *gin.Context) *utils.ReturnStatus
	LoginTOTP(email, totpCode string) (*domain.User, string, *utils.ReturnStatus)
}

type AdminService interface {
	GetSystemPolicy(ctx context.Context) (*config.SystemPolicy, *utils.ReturnStatus)
	UpdateSystemPolicy(ctx context.Context, updates map[string]any) (*config.SystemPolicy, *utils.ReturnStatus)
	CleanupExpiredFiles(ctx context.Context) (int, *utils.ReturnStatus)
}
