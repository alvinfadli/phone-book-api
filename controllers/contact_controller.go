package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "strings"
    "math"
	
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

// using query params for searching by name here
func (ctrl *ContactController) GetAll(c *gin.Context) {
    var contacts []models.Contact

    page, err := strconv.Atoi(c.Query("page"))
    if err != nil || page < 1 {
        page = 1
    }
    perPage, err := strconv.Atoi(c.Query("perPage"))
    if err != nil || perPage < 1 {
        perPage = 10
    }
    offset := (page - 1) * perPage

    name := c.Query("name")
    name = strings.ToLower(name)
    query := ctrl.DB
    if name != "" {
        query = query.Where("LOWER(name) LIKE ?", "%"+name+"%")
    }
    
    var total int64
    if err := query.Model(&models.Contact{}).Count(&total).Error; err != nil {
        c.JSON(http.StatusInternalServerError, helpers.RespondWithError(http.StatusInternalServerError, "Failed to fetch contacts"))
        return
    }
    
    if err := query.Order("name").Offset(offset).Limit(perPage).Find(&contacts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, helpers.RespondWithError(http.StatusInternalServerError, "Failed to fetch contacts"))
        return
    }

    pagination := map[string]interface{}{
        "total":    total,
        "page":     page,
        "perPage":  perPage,
        "lastPage": int(math.Ceil(float64(total) / float64(perPage))),
    }
    response := map[string]interface{}{
        "contacts":  contacts,
        "pagination": pagination,
    }
    c.JSON(http.StatusOK, helpers.RespondWithData(response))
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