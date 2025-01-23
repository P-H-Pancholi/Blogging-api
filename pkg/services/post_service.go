package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/P-H-Pancholi/Blogging-api/pkg/database"
	"github.com/P-H-Pancholi/Blogging-api/pkg/models"
)

func GetAllPosts() ([]models.Post, int, error) {
	ctx := context.Background()
	var posts []models.Post
	conn := database.Db.DbConn
	query := `SELECT
	            id,
	            title,
	            content,
	            category,
	            created_at,
	            updated_at
	        FROM posts`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return []models.Post{}, http.StatusInternalServerError, fmt.Errorf("error while running query: %w", err)
	}

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return []models.Post{}, http.StatusInternalServerError, fmt.Errorf("error while scanning output: %w", err)
		}
		tagsQuery := `SELECT tag FROM tags WHERE post_id = $1`
		tagRows, err := conn.Query(ctx, tagsQuery, post.ID)
		if err != nil {
			return []models.Post{}, http.StatusInternalServerError, fmt.Errorf("error while running query: %w", err)
		}
		for tagRows.Next() {
			var tag string
			if err := tagRows.Scan(&tag); err != nil {
				return []models.Post{}, http.StatusInternalServerError, fmt.Errorf("error while scanning output: %w", err)
			}
			post.Tags = append(post.Tags, tag)
		}
		posts = append(posts, post)
	}
	return posts, http.StatusOK, nil
}

func CreatePost(post *models.Post) (int, error) {
	ctx := context.Background()
	conn := database.Db.DbConn
	tx, err := conn.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to start transaction: %v", err)
	}

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	query := `INSERT INTO posts (title, content, category, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	if err := tx.QueryRow(ctx, query, post.Title, post.Content, post.Category, post.CreatedAt, post.UpdatedAt).Scan(&post.ID); err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to insert post: %w", err)
	}

	for _, tag := range post.Tags {
		tagQuery := "INSERT INTO tags (post_id, tag) VALUES ($1, $2)"
		if _, err := tx.Exec(ctx, tagQuery, post.ID, tag); err != nil {
			tx.Rollback(ctx)
			return http.StatusInternalServerError, fmt.Errorf("failed to insert tag: %w", err)
		}
	}
	return http.StatusOK, tx.Commit(ctx)
}

func UpdatePost(post *models.Post, id uint) (int, error) {
	ctx := context.Background()
	conn := database.Db.DbConn
	tx, err := conn.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to start transaction: %v", err)
	}

	query := `UPDATE posts 
			SET title = $1,
			content = $2,
			category = $3,
			updated_at = $4
			WHERE id = $5`

	ct, err := tx.Exec(ctx, query, post.Title, post.Content, post.Category, time.Now(), id)
	if ct.RowsAffected() == 0 {
		return http.StatusNotFound, fmt.Errorf("post not found")
	}
	if err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to update post: %w", err)
	}

	tagDeleteQuery := `DELETE FROM tags WHERE post_id = $1`

	if _, err := tx.Exec(ctx, tagDeleteQuery, id); err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to update post: %w", err)
	}

	for _, tag := range post.Tags {
		tagQuery := "INSERT INTO tags (post_id, tag) VALUES ($1, $2)"
		if _, err := tx.Exec(ctx, tagQuery, id, tag); err != nil {
			tx.Rollback(ctx)
			return http.StatusInternalServerError, fmt.Errorf("failed to update post: %w", err)
		}
	}

	getPostQuery := `SELECT id, title, content, category, created_at, updated_at FROM posts WHERE id = $1`
	row := tx.QueryRow(ctx, getPostQuery, id)
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
		tx.Rollback(ctx)
		return http.StatusInternalServerError, fmt.Errorf("failed to update post: %w", err)
	}
	return http.StatusOK, tx.Commit(ctx)
}
