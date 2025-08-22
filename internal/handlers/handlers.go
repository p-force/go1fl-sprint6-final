package handlers
<<<<<<< HEAD

import (
	"net/http"
	"os"
	"io"
	"time"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// import "encoding/json"

// Для корневого эндпоинта / нужно реализовать хендлер,
// который возвращает HTML из файла index.html
func MainHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data, err := os.ReadFile("index.html")

	if err != nil {
		http.Error(w, "Failed to load page", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := req.FormFile("myFile")

	if err != nil {
		http.Error(w, "Unable to read uploaded file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	// Передаём данные в функцию автоопределения из пакета service
	convertedString, err := service.AutoDetect(fileData)
	if err != nil {
		http.Error(w, "Conversion failed", http.StatusInternalServerError)
		return
	}

	// Создаём имя локального файла с временем и расширением
	fileName := time.Now().UTC().Format("20060102150405") + filepath.Ext(fileHeader.Filename)

	// Создаём локальный файл
	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Записываем результат конвертации в файл
	_, err = outFile.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации клиенту
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(convertedString))
}
=======
>>>>>>> 451cb3d (Initial commit)
