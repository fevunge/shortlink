// Package db
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type LinkDB struct {
	ShortURL string `json:"shorturl"`
	URL      string `json:"url"`
}

func ConnectDB() (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	sqlStmt := `CREATE TABLE IF NOT EXISTS shortlink (
		shorturl TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
		
	);`
	_, err = DB.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func InsertDB(link LinkDB, DB *sql.DB) (sql.Result, error) {
	result, err := DB.Exec("INSERT INTO shortlink(shorturl, url) VALUES(?, ?)", link.ShortURL, link.URL)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SelectDB(query string, DB *sql.DB) (LinkDB, error) {
	var shorturl, url string
	err := DB.QueryRow("SELECT shorturl, url FROM shortlink WHERE = ?", query).Scan(&shorturl, &url)
	if err != nil {
		return LinkDB{}, err
	}
	return LinkDB{
		shorturl,
		url,
	}, nil
}
