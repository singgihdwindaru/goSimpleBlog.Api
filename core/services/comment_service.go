package services

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/singgihdwindaru/goSimpleBlog.Api/core/models"
)

type CommentService struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func (service *CommentService) Create(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&comment); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := comment.CreateComment(r.Context(), service.DB, comment); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	rsp := models.CommentResponse{}
	rsp.Code = http.StatusAccepted
	rsp.Message = "Komentar anda berhasil di terima"
	respondWithJSON(w, http.StatusAccepted, rsp)
}

func (service *CommentService) Postname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postname := models.CommentsByPostname{Postname: vars["postname"]}
	rsp, err := postname.GetCommentByPostname(r.Context(), service.DB, postname)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, rsp)
}

func (service *CommentService) Guid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := models.CommentsByGuid{Guid: vars["guid"]}
	rsp, err := guid.GetCommentByGuid(r.Context(), service.DB, guid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Commentar not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, rsp)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
