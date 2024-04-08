package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"

    "phone-book-api/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // remove cors (i dont think its necessary for this project)    
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }
        c.Next()
    })

    contactController := controllers.ContactController{DB: db}

    v1 := router.Group("/api/v1")
    {
        v1.GET("/contacts", contactController.GetAll)
        v1.GET("/contacts/:id", contactController.GetByID)
        v1.POST("/contacts", contactController.Create)
		v1.POST("/contacts/search", contactController.GetByName)
        v1.PUT("/contacts/:id", contactController.Update)
        v1.DELETE("/contacts/:id", contactController.Delete)
    }

    return router
}
