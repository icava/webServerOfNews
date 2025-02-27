package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("mysql", "root@tcp(MySQL-8.2:3306)/golang")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res := db.QueryRow("SELECT id, title, full_text FROM `articles` WHERE `id` = ?", vars["id"])

	var article Article
	err = res.Scan(&article.ID, &article.Title, &article.FullText)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Статья не найдена", http.StatusNotFound)
		} else {
			handleError(w, err, http.StatusInternalServerError)
		}
		return
	}

	err = t.ExecuteTemplate(w, "show", article)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
