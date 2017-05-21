package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicRest/config"
)

var db *sql.DB

func Initialize() {
	dsn := config.GetDsn()
	db, _ = sql.Open("mysql", dsn)
}

func InsertToken(key string, value string) {
	stmt, err := db.Prepare("INSERT INTO tokens (key, value) VALUES(?, ?)")
	handleErr(err)
	_, err = stmt.Exec(key, value)
	handleErr(err)
}

func UpdateToken(key string, value string) {
	stmt, err := db.Prepare("UPDATE tokens SET value=? WHERE key=?")
	handleErr(err)
	_, err = stmt.Exec(value, key)
	handleErr(err)
}

func DeleteToken(key string) {
	stmt, err := db.Prepare("DELETE FROM tokens WHERE key=?")
	handleErr(err)
	_, err = stmt.Exec(key)
	handleErr(err)
}

func GetToken(key string) string {
	var token string
	err := db.QueryRow("SELECT value FROM tokens WHERE key=?", key).Scan(&token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
func getDbData(rows *sql.Rows) []content.Page {
	pages := []content.Page{}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var (
				id         int
				title      sql.NullString
				path       sql.NullString
				imgUrl     sql.NullString
				disqusId   sql.NullString
				act        sql.NullString
				pageNumber int
			)

			rows.Scan(
				&id,
				&title,
				&path,
				&imgUrl,
				&disqusId,
				&act,
				&pageNumber)

			pages = append(pages, content.Page{
				Id:         id,
				Title:      title.String,
				Path:       path.String,
				ImgUrl:     imgUrl.String,
				DisqusId:   disqusId.String,
				Act:        act.String,
				PageNumber: pageNumber})
		}
	}
	return pages
}
*/

/*
CREATE TABLE tokens (
	key VARCHAR(255) NOT NULL UNIQUE,
	value VARCHAR(1024) NOT NULL UNIQUE
);
*/
