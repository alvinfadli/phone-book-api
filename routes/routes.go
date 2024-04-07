package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "phone-book-api/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()
    contactController := controllers.ContactController{DB: db}

    v1 := router.Group("/api/v1")
    {
        v1.POST("/contacts", contactController.Create)
        v1.GET("/contacts", contactController.GetAll)
        v1.GET("/contacts/:id", contactController.GetByID)
        v1.PUT("/contacts/:id", contactController.Update)
        v1.DELETE("/contacts/:id", contactController.Delete)
    }

    return router
}
