package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Функция для входа
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	// Если метод POST, проверяем логин
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Открытие базы данных
		db, err := sql.Open("mysql", conToBD)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Проверка логина
		var dbPassword, uid string
		err = db.QueryRow("SELECT password, uid FROM users WHERE username = ?", username).Scan(&dbPassword, &uid)
		if err != nil || dbPassword != password {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		// Создаем сессию
		session, err := store.Get(r, "session")
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}
		session.Values["user"] = username
		session.Values["uid"] = uid // Сохраняем UID в сессию
		session.Save(r, w)

		// Перенаправляем на главную
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Отображаем форму логина
	err = t.ExecuteTemplate(w, "login", nil)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
