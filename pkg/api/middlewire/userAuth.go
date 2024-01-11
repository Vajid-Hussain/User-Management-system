package middlewire

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ClearChache(c *gin.Context) {
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Header("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	c.Next()   
}
