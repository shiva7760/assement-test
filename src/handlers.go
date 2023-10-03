package src

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDB(dbConn *gorm.DB) {
	db = dbConn
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts := []BlogPost{}
	db.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post := BlogPost{}
	db.First(&post, params["id"])
	if post.ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := BlogPost{}
	post.ID = uint(rand.Uint32())
	json.NewDecoder(r.Body).Decode(&post)
	db.Create(&post)
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post := BlogPost{}
	db.First(&post, params["id"])
	if post.ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&post)
	db.Save(&post)
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post := BlogPost{}
	db.First(&post, params["id"])
	if post.ID == 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	db.Delete(&post)
	w.WriteHeader(http.StatusNoContent)
}
