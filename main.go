package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "go-blog-author"
	dbPass = "admin1314"
	dbHost = "localhost"
	dbPort = "3306"
)

type Article struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type ResponseBody struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Article `json:"data"`
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}
