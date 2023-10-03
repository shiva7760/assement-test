package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	app "github.com/shiva7760/assement-test/src"
)

func main() {

	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&app.BlogPost{})

	router := mux.NewRouter()
	app.SetDB(db)

	router.HandleFunc("/posts", app.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", app.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", app.DeletePost).Methods("DELETE")

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
	println("container exited")
}

func connectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=assesment dbname=blogdb password=very_secret_password sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to container Retrying")
		c := 0
		for c < 12 {
			time.Sleep(2 * time.Second)
			db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=assesment dbname=blogdb password=very_secret_password sslmode=disable")
			if db != nil {
				break
			} else {
				continue
			}
		}
		if c == 12 {
			return nil, err
		}
	}
	fmt.Println("App Booted up successfully : request at http://localhost:8080/posts/")

	return db, nil
}
