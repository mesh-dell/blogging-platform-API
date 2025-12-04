package repository

import (
	"database/sql"
	"fmt"
	"time"

	model "github.com/mesh-dell/blogging-platform-API/internal/blog"
)

type BlogRepository struct {
	db *sql.DB
}

// CreateBlog implements IBlogPostsRepository.
func (r *BlogRepository) CreateBlog(req model.BlogPostRequest) (model.BlogPost, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return model.BlogPost{}, nil
	}

	res, err := tx.Exec("INSERT INTO posts (title, content, category) VALUES(?, ?, ?)",
		req.Title, req.Content, req.Category,
	)

	if err != nil {
		tx.Rollback()
		return model.BlogPost{}, fmt.Errorf("insert post: %w", err)
	}

	postID64, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return model.BlogPost{}, err
	}
	postId := int(postID64)

	for _, tag := range req.Tags {
		var tagID int
		err = tx.QueryRow("SELECT id FROM tags WHERE tag_name = ?", tag).Scan(&tagID)
		if err == sql.ErrNoRows {
			res, err := tx.Exec("INSERT INTO tags (tag_name) VALUES(?)", tag)
			if err != nil {
				tx.Rollback()
				return model.BlogPost{}, err
			}
			id64, _ := res.LastInsertId()
			tagID = int(id64)
		} else if err != nil {
			tx.Rollback()
			return model.BlogPost{}, err
		}
		_, err := tx.Exec("INSERT IGNORE INTO post_tags (post_id, tag_id) VALUES (?, ?)",
			postId, tagID,
		)
		if err != nil {
			tx.Rollback()
			return model.BlogPost{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return model.BlogPost{}, err
	}
	now := time.Now()
	return model.BlogPost{
		Id:        postId,
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
		Tags:      req.Tags,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// DeleteBlog implements IBlogPostsRepository.
func (r *BlogRepository) DeleteBlog(id int) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

// GetBlogById implements IBlogPostsRepository.
func (r *BlogRepository) GetBlogById(id int) (model.BlogPost, error) {
	var p model.BlogPost
	row := r.db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	if err := row.Scan(&p.Id, &p.Title, &p.Content, &p.Category, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return model.BlogPost{}, fmt.Errorf("post %d not found", id)
		}
		return p, err
	}
	// get tags
	rows, err := r.db.Query("SELECT t.tag_name FROM tags t JOIN post_tags pt ON t.id = pt.post_id WHERE pt.post_id = ?", id)
	if err != nil {
		return p, err
	}
	defer rows.Close()
	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return p, err
		}
		tags = append(tags, tag)
	}
	p.Tags = tags
	return p, nil
}

// GetBlogs implements IBlogPostsRepository.
func (r *BlogRepository) GetBlogs() ([]model.BlogPost, error) {
	query := `
		SELECT 
            p.id, 
            p.title, 
            p.content, 
            p.category, 
            p.created_at, 
            p.updated_at,
            t.tag_name
        FROM 
            posts p
        LEFT JOIN 
            post_tags pt ON p.id = pt.post_id
        LEFT JOIN 
            tags t ON t.id = pt.tag_id
        ORDER BY p.id	
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	postsMap := make(map[int]*model.BlogPost)
	var postOrder []int

	for rows.Next() {
		var (
			id        int
			title     string
			content   string
			category  string
			createdAt time.Time
			updatedAt time.Time
			tagName   sql.NullString
		)
		if err := rows.Scan(&id, &title, &content, &category, &createdAt, &updatedAt, &tagName); err != nil {
			return nil, err
		}
		post, exists := postsMap[id]
		if !exists {
			post = &model.BlogPost{
				Id:        id,
				Title:     title,
				Content:   content,
				Category:  category,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Tags:      []string{},
			}
			postsMap[id] = post
			postOrder = append(postOrder, id)
		}
		if tagName.Valid {
			post.Tags = append(post.Tags, tagName.String)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	posts := make([]model.BlogPost, 0, len(postOrder))
	for _, id := range postOrder {
		posts = append(posts, *postsMap[id])
	}
	return posts, nil
}

// UpdateBlog implements IBlogPostsRepository.
func (r *BlogRepository) UpdateBlog(id int, req model.BlogPostRequest) (model.BlogPost, error) {
	_, err := r.db.Exec("UPDATE posts SET title = ?, content = ?, category = ? WHERE id = ?",
		req.Title, req.Content, req.Category, id,
	)
	if err != nil {
		return model.BlogPost{}, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return model.BlogPost{}, err
	}

	if _, err := tx.Exec("DELETE FROM post_tags WHERE post_id = ?", id); err != nil {
		tx.Rollback()
		return model.BlogPost{}, err
	}

	for _, tag := range req.Tags {
		var tagId int
		if err := tx.QueryRow("SELECT id FROM tags WHERE tag_name = ?", tag).Scan(&tagId); err != nil {
			if err == sql.ErrNoRows {
				res, err := tx.Exec("INSERT INTO tags (tag_name) VALUES (?)", tag)
				if err != nil {
					tx.Rollback()
					return model.BlogPost{}, err
				}
				id64, _ := res.LastInsertId()
				tagId = int(id64)
			} else {
				tx.Rollback()
				return model.BlogPost{}, err
			}
		}
		if _, err := tx.Exec("INSERT IGNORE INTO post_tags (post_id, tag_id) VALUES (?, ?)", id, tagId); err != nil {
			tx.Rollback()
			return model.BlogPost{}, err
		}
	}
	if err := tx.Commit(); err != nil {
		return model.BlogPost{}, err
	}
	return r.GetBlogById(id)
}

func NewBlogRepository(db *sql.DB) IBlogPostsRepository {
	return &BlogRepository{db: db}
}
