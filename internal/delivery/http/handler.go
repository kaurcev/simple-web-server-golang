package http

import (
	"fmt"
	"net/http"
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

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, html404)
}
