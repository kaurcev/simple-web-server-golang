package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const html404 = `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>404</title>
</head>
<body>
    <h1>404</h1>
    <hr />
    <p>Данный ресурс не найден или не существовал вовсе</p>
</body>
</html>`

func customFSWrap(publicDir string, errorHandler http.HandlerFunc) http.Handler {
	fs := http.FileServer(http.Dir(publicDir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := publicDir + r.URL.Path
		stat, err := os.Stat(path)

		if os.IsNotExist(err) || (err == nil && stat.IsDir()) {
			if err == nil && stat.IsDir() {
				if _, indexErr := os.Stat(path + "/index.html"); indexErr == nil {
					fs.ServeHTTP(w, r)
					return
				}
			}

			errorHandler(w, r)
			return
		}

		fs.ServeHTTP(w, r)
	})
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, html404)
}

func main() {
	wrappedFS := customFSWrap("./public", handle404)

	http.Handle("/", wrappedFS)

	log.Println("Сервер запущен на http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
