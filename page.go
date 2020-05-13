package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
	})
}