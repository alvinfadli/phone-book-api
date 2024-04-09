package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
	
	"phone-book-api/models"
	"phone-book-api/helpers"
)

type ContactController struct {
    DB *gorm.DB
}

func (ctrl *ContactController) Create(c *gin.Context) {
    var contact models.Contact
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, helpers.RespondWithError(http.StatusBadRequest, err.Error()))
        return
    }

	// check if name and phone is empty
	if contact.Name == "" || contact.Phone == "" {
        c.JSON(http.StatusBadRequest, helpers.RespondWithError(http.StatusBadRequest, "Name and Phone are required"))
        return
    }

    ctrl.DB.Create(&contact)
    c.JSON(http.StatusCreated, helpers.RespondWithData(contact))
}

func (ctrl *ContactController) GetAll(c *gin.Context) {
    var contacts []models.Contact
    if err := ctrl.DB.Order("name").Find(&contacts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, helpers.RespondWithError(http.StatusInternalServerError, "Failed to fetch contacts"))
        return
    }
    c.JSON(http.StatusOK, helpers.RespondWithData(contacts))
}

func (ctrl *ContactController) GetByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var contact models.Contact
    if err := ctrl.DB.First(&contact, id).Error; err != nil {
        c.JSON(http.StatusNotFound, helpers.RespondWithError(http.StatusNotFound, "Record not found!"))
        return
    }
    c.JSON(http.StatusOK, helpers.RespondWithData(contact))
}

func (ctrl *ContactController) GetByName(c *gin.Context) {
	var req struct {
        Name string `json:"name" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, helpers.RespondWithError(http.StatusBadRequest, err.Error()))
        return
    }

    var contacts []models.Contact
    if err := ctrl.DB.Where("name ILIKE ?", "%"+req.Name+"%").Find(&contacts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, helpers.RespondWithError(http.StatusInternalServerError, "Failed to search contacts"))
        return
    }

    c.JSON(http.StatusOK, helpers.RespondWithData(contacts))
}

func (ctrl *ContactController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var contact models.Contact
    if err := ctrl.DB.First(&contact, id).Error; err != nil {
        c.JSON(http.StatusNotFound, helpers.RespondWithError(http.StatusNotFound, "Record not found!"))
        return
    }

    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, helpers.RespondWithError(http.StatusBadRequest, err.Error()))
        return
    }

    ctrl.DB.Save(&contact)
    c.JSON(http.StatusOK, helpers.RespondWithData(contact))
}

func (ctrl *ContactController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var contact models.Contact
    if err := ctrl.DB.First(&contact, id).Error; err != nil {
        c.JSON(http.StatusNotFound, helpers.RespondWithError(http.StatusNotFound, "Record not found!"))
        return
    }

    ctrl.DB.Delete(&contact)
    c.JSON(http.StatusOK, helpers.RespondWithData(nil))
}