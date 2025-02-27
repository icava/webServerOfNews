package handlers

import (
	"html/template"
	"net/http"
)

func Contacts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "contacts", nil)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
