package bootstrap

import (
	"time"

	"github.com/gin-contrib/cors"
)

func NewCorsConfig(env *Env) cors.Config {
	corsConf := cors.Config{
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}

	// TODO: split dev env and prod env
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host"}

	return corsConf
}
