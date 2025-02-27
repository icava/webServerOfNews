package handlers

import (
	"database/sql"
	"regexp"

	"html/template"

	"net/http"
	"net/mail"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid" // Пакет для генерации уникальных идентификаторов
)

// Функция для отображения страницы регистрации
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	// Парсинг шаблона регистрации
	t, err := template.ParseFiles("templates/register.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Рендерим страницу регистрации
	err = t.ExecuteTemplate(w, "register", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isValidUsername(username string) bool {
	pattern := regexp.MustCompile(`[a-zA-Z0-9_]`)
	return pattern.MatchString(username)
}

func isValidPassword(password, confirmPassword string) bool {
	pattern := regexp.MustCompile(`[a-zA-Z0-9]`)
	return pattern.MatchString(password) && password == confirmPassword
}

func isValidEmailReg(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Функция для обработки регистрации нового пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем данные из формы
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm-password")

		if isValidPassword(password, confirmPassword) && isValidEmailReg(email) && isValidUsername(username) {
			// Генерация уникального идентификатора (UUID)
			uid := uuid.New().String()

			// Открытие базы данных
			db, err := sql.Open("mysql", conToBD)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer db.Close()
			if db.QueryRow("SELECT username, email FROM users WHERE username =? OR email =?", username, email).Scan() != sql.ErrNoRows {
				http.Error(w, "Такой логин или email уже зарегистрирован", http.StatusBadRequest)
				return
			} else {
				// Добавляем пользователя в базу данных с кодом и временем его создания
				_, err = db.Exec("INSERT INTO users (uid, username, email, password) VALUES (?, ?, ?, ?)", uid, username, email, password)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		} else {
			http.Error(w, "Неверный формат логина или пароля", http.StatusBadRequest)
			return
		}

	}
}
