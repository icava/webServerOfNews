package handlers

import (
	"fmt"
	"net/http"
)

// Вспомогательная функция для обработки ошибок
func handleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		fmt.Println("Ошибка:", err)
	}
}
