package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func index(c *gin.Context) {
	if curPage, err := strconv.Atoi(c.DefaultQuery("page", "1")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, c.MustGet("Repository").(*Repository).all(curPage))
	}

}

func create(c *gin.Context) {
	var json BlogPost

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if c.MustGet("Repository").(*Repository).add(&json) == true {
		c.JSON(http.StatusCreated, json)
	}
}

func show(c *gin.Context) {
	err, result := c.MustGet("Repository").(*Repository).get(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, result)
}

func remove(c *gin.Context) {
	if err := c.MustGet("Repository").(*Repository).delete(c.Param("id")); err != nil {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
