package filesystem

import (
	"net/http"
	"os"
)

type FileSystemService struct {
	publicDir string
}

func NewFileSystemService(publicDir string) *FileSystemService {
	return &FileSystemService{publicDir: publicDir}
}

func (s *FileSystemService) WrapHandler(errorHandler http.HandlerFunc) http.Handler {
	fs := http.FileServer(http.Dir(s.publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := s.publicDir + r.URL.Path
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
