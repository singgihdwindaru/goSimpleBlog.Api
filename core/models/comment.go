package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type Comment struct {
	FirtsName string
	LastName  string
	Commentar string
	Parent    int
	PostName  string
}
type CommentsByPostname struct {
	Postname string
}
type CommentsByGuid struct {
	Guid string
}

type CommentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type PostnameResponse struct {
	FirtsName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Commentar string `json:"commentar"`
	Guid      string `json:"guid"`
	Status    bool   `json:"status"`
}

func (comment *Comment) CreateComment(ctx context.Context, db *sql.DB, data Comment) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	query := "insert into comments(firstname,lastname,commentar,postname) values(?,?,?,?)"
	_, err = tx.ExecContext(ctx, query, data.FirtsName, data.LastName, data.Commentar, data.PostName)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	return err
}
func (postname *CommentsByPostname) GetCommentByPostname(ctx context.Context, db *sql.DB, data CommentsByPostname) ([]PostnameResponse, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	query := "SELECT firstname, lastname, commentar,'' AS guid,"
	query += "case WHEN pub_date is NULL then false else true end as status "
	query += "FROM comments "
	query += "WHERE postname=?"
	rows, err := tx.QueryContext(ctx, query, postname.Postname)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	response := []PostnameResponse{}
	for rows.Next() {
		var rsp PostnameResponse
		err = rows.Scan(&rsp.FirtsName, &rsp.LastName, &rsp.Commentar, &rsp.Guid, &rsp.Status)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, rsp)
	}
	return response, nil
}

func (comment *CommentsByGuid) GetCommentByGuid(ctx context.Context, db *sql.DB, guid CommentsByGuid) (PostnameResponse, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	query := "SELECT firstname, lastname, commentar,'' AS guid,"
	query += "case WHEN pub_date is NULL then false else true end as status "
	query += "FROM comments "
	query += "WHERE guid=?"
	rows, err := tx.QueryContext(ctx, query, guid.Guid)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	rsp := PostnameResponse{}
	if rows.Next() {
		err = rows.Scan(&rsp.FirtsName, &rsp.LastName, &rsp.Commentar, &rsp.Guid, &rsp.Status)
		if err != nil {
			log.Fatal(err)
		}
		return rsp, nil
	} else {
		return rsp, errors.New("There is no comment")
	}
}
