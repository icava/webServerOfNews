package handlers

import (
	"net/http"
)

// Функция для выхода из аккаунта
func Logout(w http.ResponseWriter, r *http.Request) {
	// Получаем сессию
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Удаляем данные пользователя из сессии
	session.Values["user"] = nil
	session.Values["uid"] = nil

	// Сохраняем сессию (обновляем её)
	session.Save(r, w)

	// Перенаправляем на страницу входа
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
