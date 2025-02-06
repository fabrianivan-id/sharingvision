package repositories

import (
    "database/sql"
    "fmt"
    "your_project/models"
)

type ArticleRepository struct {
    db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
    return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *models.Article) error {
    query := `INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)`
    res, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status)
    if err != nil {
        return err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    article.ID = int(id)
    return nil
}

func (r *ArticleRepository) GetByID(id int) (*models.Article, error) {
    query := `SELECT id, title, content, category, status FROM posts WHERE id = ?`
    row := r.db.QueryRow(query, id)
    article := &models.Article{}
    err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Status)
    if err != nil {
        return nil, err
    }
    return article, nil
}

func (r *ArticleRepository) List(limit, offset int) ([]models.Article, error) {
    query := `SELECT id, title, content, category, status FROM posts LIMIT ? OFFSET ?`
    rows, err := r.db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []models.Article
    for rows.Next() {
        var a models.Article
        err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Category, &a.Status)
        if err != nil {
            return nil, err
        }
        articles = append(articles, a)
    }
    return articles, nil
}

func (r *ArticleRepository) Update(id int, article *models.Article) error {
    query := `UPDATE posts SET title=?, content=?, category=?, status=? WHERE id=?`
    _, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status, id)
    return err
}

func (r *ArticleRepository) Delete(id int) error {
    query := `DELETE FROM posts WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}