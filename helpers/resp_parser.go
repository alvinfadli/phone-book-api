package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondWithData(data interface{}) gin.H {
    return gin.H{
        "data":   data,
        "status": gin.H{"code": http.StatusOK, "message": "OK"},
    }
}

func RespondWithError(code int, message string) gin.H {
    return gin.H{
        "error":  message,
        "status": gin.H{"code": code, "message": http.StatusText(code)},
    }
}