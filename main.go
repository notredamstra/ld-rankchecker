package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	httpClient *http.Client
	baseURL    string
)

func main() {
	httpClient = &http.Client{}
	baseURL = "https://api.ldjam.com"

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob(getEnv("TEMPLATE_PATTERN", "./public/templates/*.html"))
	r.GET("/", showPage)
	r.POST("/", getRank)
	r.Static("/assets", getEnv("STATIC_PATTERN", "./public"))
	err := r.Run(":" + getEnv("PORT", "8080"))
	if err != nil {
		panic(err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
