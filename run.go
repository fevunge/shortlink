// package main
package main

import (
	"net/http"
	"time"

	"shortlink/api"
	"shortlink/db"
)

func Run() error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	handler := api.NewHandler(db)
	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 1,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
