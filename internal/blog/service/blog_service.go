package service

import (
	model "github.com/mesh-dell/blogging-platform-API/internal/blog"
	"github.com/mesh-dell/blogging-platform-API/internal/blog/repository"
)

type BlogPostService struct {
	repo repository.IBlogPostsRepository
}

func NewBlogPostService(r repository.IBlogPostsRepository) *BlogPostService {
	return &BlogPostService{repo: r}
}

func (s *BlogPostService) CreateBlog(req model.BlogPostRequest) (model.BlogPost, error) {
	return s.repo.CreateBlog(req)
}

func (s *BlogPostService) GetBlogById(id int) (model.BlogPost, error) {
	return s.repo.GetBlogById(id)
}

func (s *BlogPostService) GetBlogs(q string) ([]model.BlogPost, error) {
	return s.repo.GetBlogs(q)
}

func (s *BlogPostService) UpdateBlog(id int, req model.BlogPostRequest) (model.BlogPost, error) {
	return s.repo.UpdateBlog(id, req)
}

func (s *BlogPostService) DeleteBlog(id int) error {
	return s.repo.DeleteBlog(id)
}
