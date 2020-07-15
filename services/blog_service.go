package services

import (
	blogs "github.com/kadekchrisna/nothoax-blog-api/domains/blog"
	utils "github.com/kadekchrisna/nothoax-blog-api/utils/errors"
)

var (
	BlogServices blogServiceInterface = &blogServiceStruct{}
)

type (
	blogServiceInterface interface {
		GetBlogByID(int64) (*blogs.Blog, utils.ResErr)
		GetAllBlogsByStatusActive(bool) (*blogs.Blogs, utils.ResErr)
	}
	blogServiceStruct struct {
	}
)

func (bs *blogServiceStruct) GetBlogByID(ID int64) (*blogs.Blog, utils.ResErr) {
	var blog blogs.Blog
	blog.Id = ID
	if err := blog.GetBlogByID(); err != nil {
		return nil, err
	}

	return &blog, nil
}

func (bs *blogServiceStruct) GetAllBlogsByStatusActive(active bool) (*blogs.Blogs, utils.ResErr) {
	var blog blogs.Blog
	blog.IsActive = 1
	if !active {
		blog.IsActive = 0
	}
	result, err := blog.GetAllBlogs()
	if err != nil {
		return nil, err
	}
	return &result, nil
}
