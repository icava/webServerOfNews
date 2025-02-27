package handlers

import (
	"database/sql"
	"net/http"

	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func isValidEmail(email string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return pattern.MatchString(email)
}

// Функция для обновления данных профиля
func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	// Проверяем авторизацию
	authorized, uid := checkAuth(r)
	if !authorized {
		// Если не авторизован, перенаправляем на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// Получаем новые данные из формы
		newEmail := r.FormValue("email")
		newPassword := r.FormValue("password")
		confirmPassword := r.FormValue("confirm-password")

		if !isValidEmail(newEmail) {
			http.Error(w, "Неверный формат email", http.StatusBadRequest)
			return
		}
		// Проверка совпадения паролей
		if newPassword != confirmPassword {
			http.Error(w, "Пароли не совпадают", http.StatusBadRequest)
			return
		}

		// Открытие базы данных
		db, err := sql.Open("mysql", "root@tcp(MySQL-8.2:3306)/golang")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Если пароль был введён, обновляем его, иначе — оставляем прежним
		var updateQuery string
		var args []interface{}

		if newPassword != "" {
			updateQuery = "UPDATE `users` SET `email` = ?, `password` = ? WHERE `uid` = ?"
			args = append(args, newEmail, newPassword, uid)
		} else {
			updateQuery = "UPDATE `users` SET `email` = ? WHERE `uid` = ?"
			args = append(args, newEmail, uid)
		}

		// Выполнение запроса на обновление данных пользователя
		_, err = db.Exec(updateQuery, args...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправление на страницу аккаунта
		http.Redirect(w, r, "/account", http.StatusSeeOther)
	}
}
