package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
)

// Функция для отображения страницы аккаунта
func AccountPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем авторизацию и извлекаем uid из сессии
	authorized, uid := checkAuth(r)
	if !authorized {
		// Если не авторизован, перенаправляем на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Открытие базы данных
	db, err := sql.Open("mysql", "root@tcp(MySQL-8.2:3306)/golang")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Запрос к базе данных для получения данных пользователя по UID
	var username, email string
	err = db.QueryRow("SELECT username, email FROM users WHERE uid = ?", uid).Scan(&username, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Пользователь не найден", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Парсинг шаблона
	t, err := template.ParseFiles("templates/account.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем структуру для передачи данных в шаблон
	user := struct {
		Username string
		Email    string
		UID      string
	}{
		Username: username,
		Email:    email,
		UID:      uid,
	}

	// Рендерим страницу аккаунта с данными пользователя
	err = t.ExecuteTemplate(w, "account", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
