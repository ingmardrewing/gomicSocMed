package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ingmardrewing/gomicSocMed/config"
)

/*
CREATE TABLE tokens (
	socmed_id VARCHAR(255),
	tkey VARCHAR(255) NOT NULL,
	tvalue VARCHAR(1024) ,
	CONSTRAINT UC_key UNIQUE( tkey )
);
*/
var db *sql.DB

func Initialize() {
	log.Println("initializing db")
	dsn := config.GetDsn()
	db_local, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("DB error while connecting to db server: ", err)
	}
	if db_local == nil {
		log.Println("db_local is nil")
	}
	db = db_local
}

func Check() error {
	if db == nil {
		return errors.New("db is nil")
	}
	return nil
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
	log.Println("getting token for", key)
	var token string
	log.Println("getting token ..1")
	if db == nil {
		log.Println("db is nil")
	}
	scanner := db.QueryRow("SELECT tvalue FROM tokens WHERE tkey=?", key)
	log.Println("getting token ..2")
	err := scanner.Scan(&token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("found token", token)
	return token
}

func GetId(key string) string {

	var socmed_id string
	err := db.QueryRow("SELECT socmed_id FROM tokens WHERE tkey=?", key).Scan(&socmed_id)
	if err != nil {
		log.Fatal(err)
	}
	return socmed_id
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
