package controller

import (
	"apigee-portal/v2/bootstrap"
	"apigee-portal/v2/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
	Env           *bootstrap.Env
}

// @Summary User Logout
// @Description Default user login
// @Tags Logout
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /logout [get]
func (lc *LogoutController) Logout(c *gin.Context) {
	clearCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(c.Writer, &clearCookie)
	c.JSON(http.StatusOK, gin.H{"message": "Logout success."})
}
