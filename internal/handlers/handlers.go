package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	original := string(content)

	converted, err := service.ProcessData(original)
	if err != nil {
		http.Error(w, "Ошибка при конвертации", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("20060102_150405")
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}

	outputFilename := "converted_" + timestamp + ext

	err = os.WriteFile(outputFilename, []byte(converted), 0644)
	if err != nil {
		http.Error(w, "Ошибка при сохранении результата", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(original))
	w.Write([]byte("\n\n"))

	w.Write([]byte("Сконвертированный текст:\n"))
	w.Write([]byte(converted))
	w.Write([]byte("\n\nРезультат сохранен в файл: "))
	w.Write([]byte(outputFilename))
}
