package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/singgihdwindaru/goSimpleBlog.Api/core/services"
)

type App struct {
	Router *mux.Router
	Db     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, password, dbname)

	var err error
	a.Db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	commentService := services.CommentService{
		DB:       a.Db,
		Validate: &validator.Validate{},
	}
	a.CommentRouting(commentService)
}

func (a *App) Run(addr string) {
	// Use default options
	handler := cors.Default().Handler(a.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
