package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "your_project/models"
    "your_project/repositories"

    "github.com/go-playground/validator/v10"
    "github.com/gorilla/mux"
)

func CreateArticleHandler(repo *repositories.ArticleRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req models.ArticleRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        validate := validator.New()
        if err := validate.Struct(req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        article := models.Article{
            Title:    req.Title,
            Content:  req.Content,
            Category: req.Category,
            Status:   req.Status,
        }

        if err := repo.Create(&article); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(struct{}{})
    }
}

func GetArticlesHandler(repo *repositories.ArticleRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        limit, err := strconv.Atoi(vars["limit"])
        if err != nil {
            http.Error(w, "Invalid limit", http.StatusBadRequest)
            return
        }
        offset, err := strconv.Atoi(vars["offset"])
        if err != nil {
            http.Error(w, "Invalid offset", http.StatusBadRequest)
            return
        }

        articles, err := repo.List(limit, offset)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        response := make([]models.ArticleResponse, len(articles))
        for i, a := range articles {
            response[i] = models.ArticleResponse{
                Title:    a.Title,
                Content:  a.Content,
                Category: a.Category,
                Status:   a.Status,
            }
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func GetArticleHandler(repo *repositories.ArticleRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        article, err := repo.GetByID(id)
        if err != nil {
            http.Error(w, "Article not found", http.StatusNotFound)
            return
        }

        response := models.ArticleResponse{
            Title:    article.Title,
            Content:  article.Content,
            Category: article.Category,
            Status:   article.Status,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func UpdateArticleHandler(repo *repositories.ArticleRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        var req models.ArticleRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        validate := validator.New()
        if err := validate.Struct(req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        article := models.Article{
            Title:    req.Title,
            Content:  req.Content,
            Category: req.Category,
            Status:   req.Status,
        }

        if err := repo.Update(id, &article); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(struct{}{})
    }
}

func DeleteArticleHandler(repo *repositories.ArticleRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "Invalid ID", http.StatusBadRequest)
            return
        }

        if err := repo.Delete(id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(struct{}{})
    }
}