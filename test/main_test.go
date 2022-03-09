package test

import (
	"log"
	"os"
	"testing"

	"github.com/singgihdwindaru/goSimpleBlog.Api/core/app"
)

var a app.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME_TEST"),
		os.Getenv("APP_DB_PASSWORD_TEST"),
		os.Getenv("APP_DB_NAME_TEST"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.Db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.Db.Exec("DELETE FROM comments")
	a.Db.Exec("ALTER SEQUENCE comment_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS comments
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	article_id bigint(20) NOT NULL,
	name varchar(50) NOT NULL COMMENT 'nama pengkomentar',
	email varchar(50) NOT NULL COMMENT 'email pengkomentar',
	is_author tinyint(1) NOT NULL COMMENT 'apakah author yang menulis atau menjawab komentar',
	pub_date datetime NOT NULL COMMENT 'tanggal publikasi komentar',
	updated_date datetime DEFAULT NULL COMMENT 'tanggal komentar di update setelah publikasi pertama kali',
	PRIMARY KEY (id),
	UNIQUE KEY comment_id_seq (id) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8`
