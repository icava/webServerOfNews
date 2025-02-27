package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	full_text := r.FormValue("content")
	bol, uid := checkAuth(r)

	db, err := sql.Open("mysql", "root@tcp(MySQL-8.2:3306)/golang")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer db.Close()
	if bol {
		insert, err := db.Prepare("INSERT INTO `articles` (`title`, `full_text`, `uid`) VALUES (?, ?, ?)")
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
		defer insert.Close()

		_, err = insert.Exec(title, full_text, uid)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Create(w http.ResponseWriter, r *http.Request) {
	authorized, _ := checkAuth(r)

	if !authorized {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "create", nil)
}
