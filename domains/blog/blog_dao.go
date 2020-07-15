package blogs

import (
	"fmt"
	"strings"

	"github.com/kadekchrisna/nothoax-blog-api/datasources"
	utils "github.com/kadekchrisna/nothoax-blog-api/utils/errors"
)

// GetAllBlogs querying to retreive all blogs
func (b *Blog) GetAllBlogs() (Blogs, utils.ResErr) {
	connection, errConnection := datasources.DBConnection.OpenConnection()
	if errConnection != nil {
		return []Blog{}, errConnection
	}

	stmn, errStmn := connection.Prepare("SELECT id_blog, blog_title, blog_body, is_active, created_at FROM blog WHERE is_active=? LIMIT 10;")
	if errStmn != nil {
		return []Blog{}, utils.NewInternalServerError("Error when preparing statement", errStmn)
	}
	defer stmn.Close()

	result, errQuery := stmn.Query(b.IsActive)
	if errQuery != nil {
		return []Blog{}, utils.NewInternalServerError("Error when preparing statement", errStmn)
	}
	defer result.Close()

	blogs := make(Blogs, 0)
	for result.Next() {
		var blog Blog
		if err := result.Scan(&blog.Id, &blog.Title, &blog.Body, &blog.IsActive, &blog.CreatedAt); err != nil {
			return []Blog{}, utils.NewInternalServerError("Error when parsing blogs", err)
		}
		blogs = append(blogs, blog)
	}

	if len(blogs) == 0 {
		return []Blog{}, utils.NewNotFoundError("Not a single blog found.")
	}
	return blogs, nil
}

// GetBlogByID getting blog by id_blog
func (b *Blog) GetBlogByID() utils.ResErr {
	connection, errConnection := datasources.DBConnection.OpenConnection()
	if errConnection != nil {
		return errConnection
	}

	stmn, errStmn := connection.Prepare("SELECT id_blog, blog_title, blog_body, is_active, created_at FROM blog WHERE id_blog=?;")
	if errStmn != nil {
		return utils.NewInternalServerError("Error when preparing statement", errStmn)
	}
	defer stmn.Close()

	result := stmn.QueryRow(b.Id)
	if err := result.Scan(&b.Id, &b.Title, &b.Body, &b.IsActive, &b.CreatedAt); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "no rows") {
			return utils.NewNotFoundError(fmt.Sprintf("blogs not found with id %v", b.Id))
		}
		return utils.NewInternalServerError("Error when parsing blogs", err)
	}
	return nil
}

// func (b *blog) CreateBlog() utils.ResErr {
// 	connection, errConnection := datasources.DBConnection.OpenConnection()
// 	if errConnection != nil {
// 		return errConnection
// 	}
// 	connection.BeginTx()
// }
