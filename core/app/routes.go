package app

import "github.com/singgihdwindaru/goSimpleBlog.Api/core/services"

func (app *App) CommentRouting(service services.CommentService) {
	app.Router.HandleFunc("/", service.Create).Methods("POST")
	app.Router.HandleFunc("/article/{postname}", service.Postname).Methods("GET")
	app.Router.HandleFunc("/{guid}}", service.Guid).Methods("GET")
}
