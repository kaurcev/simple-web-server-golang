package main

import (
	"log"
	appHttp "my-server/internal/delivery/http"
	"my-server/internal/delivery/http/middleware"
	"my-server/internal/filesystem"
	"net/http"
)

func main() {
	fsService := filesystem.NewFileSystemService("./public")
	wrappedFS := fsService.WrapHandler(appHttp.Handle404)
	handlerWithLogging := middleware.Logging(wrappedFS)
	http.Handle("/", handlerWithLogging)

	log.Println("Сервер запущен на http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
