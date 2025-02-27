package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Структура для статьи
type Article struct {
	ID       uint16 `json:"id"`
	Title    string `json:"title"`
	FullText string `json:"full_text"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError) // Использование handleError
		return
	}

	db, err := sql.Open("mysql", conToBD)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res, err := db.Query("SELECT id, title, full_text FROM `articles`")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer res.Close()

	var posts []Article
	for res.Next() {
		var post Article
		err = res.Scan(&post.ID, &post.Title, &post.FullText)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	err = t.ExecuteTemplate(w, "index", posts)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
