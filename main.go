package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)

	log.Println("Сервер запущен на http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
