package handlers

import (
	"bytes"
	"go1fl-sprint6-final-tpl/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("go1fl-sprint6-final-tpl/go1fl-sprint6-final-tpl/index.html")
	if err != nil {
		http.Error(w, "ошибка при загрузке страницы", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "не удалось распарсить форму", http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)

	if err != nil {
		http.Error(w, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	line, err := service.Convert(buf.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nameLocalFile := time.Now().UTC().Format("20060102_150405") + filepath.Ext(handler.Filename)
	localFile, err := os.OpenFile(nameLocalFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer localFile.Close()

	_, err = localFile.Write([]byte(line))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(line))
}