package routes

import (
	"github.com/gin-gonic/gin"
	blog "github.com/kadekchrisna/nothoax-blog-api/controllers/blog"
	ping "github.com/kadekchrisna/nothoax-blog-api/controllers/ping"
)

var (
	// Router for API
	router = gin.Default()
)

// Routes lists endpoints
func Routes() *gin.Engine {
	router.GET("/ping", ping.PingController.Ping)

	router.GET("/blog/get/:blog_id", blog.BlogController.GetBlog)
	router.GET("/blog/all/:active", blog.BlogController.GetBlogs)

	return router
}
