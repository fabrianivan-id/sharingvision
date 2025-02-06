package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"
    "your_project/handlers"
    "your_project/models"
    "your_project/repositories"

    "github.com/gorilla/mux"
    _ "github.com/joho/godotenv/autoload"
    _ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    dsn := os.Getenv("DB_DSN") // e.g., "user:password@tcp(localhost:3306)/article"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    runMigrations(dsn)

    repo := repositories.NewArticleRepository(db)

    router := mux.NewRouter()
    router.HandleFunc("/article", handlers.CreateArticleHandler(repo)).Methods("POST")
    router.HandleFunc("/article/{limit}/{offset}", handlers.GetArticlesHandler(repo)).Methods("GET")
    router.HandleFunc("/article/{id}", handlers.GetArticleHandler(repo)).Methods("GET")
    router.HandleFunc("/article/{id}", handlers.UpdateArticleHandler(repo)).Methods("PUT", "POST", "PATCH")
    router.HandleFunc("/article/{id}", handlers.DeleteArticleHandler(repo)).Methods("DELETE", "POST")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func runMigrations(dsn string) {
    m, err := migrate.New("file://migrations", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }
}