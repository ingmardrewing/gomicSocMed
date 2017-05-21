package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicSocMed/config"
)

var db *sql.DB

func Initialize() {
	dsn := config.GetDsn()
	db, _ = sql.Open("mysql", dsn)
}

func InsertToken(key string, value string) {
	stmt, err := db.Prepare("INSERT INTO tokens (tkey, tvalue) VALUES(?, ?)")
	handleErr(err)
	_, err = stmt.Exec(key, value)
	handleErr(err)
}

func UpdateToken(key string, value string) {
	stmt, err := db.Prepare("UPDATE tokens SET tvalue=? WHERE tkey=?")
	handleErr(err)
	_, err = stmt.Exec(value, key)
	handleErr(err)
}

func DeleteToken(key string) {
	stmt, err := db.Prepare("DELETE FROM tokens WHERE tkey=?")
	handleErr(err)
	_, err = stmt.Exec(key)
	handleErr(err)
}

func GetToken(key string) string {
	var token string
	err := db.QueryRow("SELECT tvalue FROM tokens WHERE tkey=?", key).Scan(&token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func TokenExists(key string) bool {
	var amount string
	err := db.QueryRow("SELECT count(*) FROM tokens WHERE tkey=?", key).Scan(&amount)
	if err != nil {
		log.Fatal(err)
	}
	return amount == "1"
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
	tkey VARCHAR(255) NOT NULL,
	tvalue VARCHAR(1024) ,
	CONSTRAINT UC_key UNIQUE( tkey )
);
*/
