package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/builder"
	db2 "monopc-starter/internal/db"
	"net/http"
)

func main() {
	cfg := config.LoadConfig(".env")
	db, err := db2.NewPostgresGormDB(&cfg.Database)

	if err != nil {
		panic(err)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	BuildHandler(cfg, router, db)

	if err := router.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func BuildHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	builder.BuildAuthHandler(cfg, router, db)
	builder.BuildUserHandler(cfg, router, db)
}
