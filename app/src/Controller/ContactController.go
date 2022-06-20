package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nochat/app/src/Database"
	"nochat/app/src/Helper"
	"nochat/app/src/Repository"
	"nochat/app/src/Request"
)

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"upload": "upload"})
}

func Search(c *gin.Context) {
	var filter Request.SearchByPhoneFilter

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(Helper.FormattingPhone(filter.Phone))
}

func SearchAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"searchAll": "searchAll"})
}

func Count(c *gin.Context) {
	count, err := Repository.GetCountRecords()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"value": count})
	}
}

func Last(c *gin.Context) {
	var contacts []Database.Contact

	err := Repository.GetLimitedAmount(&contacts, 500)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contacts)
	}
}
