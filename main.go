package main

import (
	"fmt"
	"log"
	"mymodule/handlers" // Импортируем пакет handlers
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunc() {
	rtr := mux.NewRouter()

	// Регистрация маршрутов
	rtr.HandleFunc("/", handlers.Index).Methods("GET")
	rtr.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
	rtr.HandleFunc("/register", handlers.RegisterPage).Methods("GET")
	rtr.HandleFunc("/register", handlers.RegisterUser).Methods("POST")

	rtr.HandleFunc("/account", handlers.AccountPage).Methods("GET")
	rtr.HandleFunc("/account/update", handlers.UpdateAccount).Methods("POST")
	rtr.HandleFunc("/logout", handlers.Logout).Methods("GET")
	rtr.HandleFunc("/create", handlers.Create).Methods("GET")
	rtr.HandleFunc("/contacts", handlers.Contacts).Methods("GET")
	rtr.HandleFunc("/save_article", handlers.SaveArticle).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", handlers.ShowArticle).Methods("GET")
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", rtr)

	fmt.Println("Сервер запущен")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleFunc()
}

/*
Планы на вечер:
1. Вывод никнейма с ссылкой на профиль
2. Градиент блока для постов
3. Вывод даты публикации
4. Сортировка постов
5. Вывод комментариев
*/
