package repository

import model "github.com/mesh-dell/blogging-platform-API/internal/blog"

type IBlogPostsRepository interface {
	CreateBlog(req model.BlogPostRequest) (model.BlogPost, error)
	GetBlogById(id int) (model.BlogPost, error)
	GetBlogs() ([]model.BlogPost, error)
	UpdateBlog(id int, req model.BlogPostRequest) (model.BlogPost, error)
	DeleteBlog(id int) error
}
