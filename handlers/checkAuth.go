package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Создаем store для работы с сессиями
var store = sessions.NewCookieStore([]byte("secret-key"))

// Функция для проверки авторизации// Функция для проверки авторизации
func checkAuth(r *http.Request) (bool, string) {
	session, err := store.Get(r, "session")
	if err != nil {
		return false, ""
	}

	// Проверка наличия данных в сессии
	uid, ok := session.Values["uid"]
	if !ok || uid == nil {
		return false, ""
	}

	// Преобразование в строку
	uidStr, ok := uid.(string)
	if !ok {
		return false, ""
	}

	return true, uidStr
}
