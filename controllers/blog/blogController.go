package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kadekchrisna/nothoax-blog-api/services"
	utils "github.com/kadekchrisna/nothoax-blog-api/utils/errors"
)

var (
	BlogController blogControllerInterface = &blogController{}
)

type blogControllerInterface interface {
	GetBlogs(*gin.Context)
	GetBlog(*gin.Context)
}

type blogController struct {
}

func parseIDParam(ID string) (int64, utils.ResErr) {
	blogID, errParam := strconv.ParseInt(ID, 10, 64)
	if errParam != nil {
		error := utils.NewBadRequestError("invalid id format")
		return 0, error
	}
	return blogID, nil
}

func parseActiveParam(active string) (bool, utils.ResErr) {
	blogID, errParam := strconv.ParseBool(active)
	if errParam != nil {
		error := utils.NewBadRequestError("invalid id format")
		return false, error
	}
	return blogID, nil
}

// GetBlogs for getting all active blogs
func (ct *blogController) GetBlogs(c *gin.Context) {
	blogActive, errParse := parseActiveParam(c.Param("active"))
	if errParse != nil {
		c.JSON(errParse.Status(), errParse)
		return
	}
	result, errFetch := services.BlogServices.GetAllBlogsByStatusActive(blogActive)
	if errFetch != nil {
		c.JSON(errFetch.Status(), errFetch)
		return
	}
	c.JSON(http.StatusFound, result)
	return
}

// Getting blog by id_blog
func (ct *blogController) GetBlog(c *gin.Context) {
	blogID, errParse := parseIDParam(c.Param("blog_id"))
	if errParse != nil {
		c.JSON(errParse.Status(), errParse)
		return
	}

	result, errFetch := services.BlogServices.GetBlogByID(blogID)
	if errFetch != nil {
		c.JSON(errFetch.Status(), errFetch)
		return
	}

	c.JSON(http.StatusFound, result)
	return
}
