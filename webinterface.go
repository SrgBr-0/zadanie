package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"
)

// refreshFunction обновляет значение, например, считывая из базы данных
func refreshFunction() (interface{}, error) {
	// Здесь может быть логика обновления значения, например, запрос к БД
	return "example_value", nil
}

// CacheService представляет сервис с кэшированием в памяти
type CacheService struct {
	cache     map[string]interface{}
	mutex     sync.RWMutex
	expiry    time.Duration
	refreshFn func() (interface{}, error)
}

// NewCacheService создает новый экземпляр CacheService
func NewCacheService(expiry time.Duration, refreshFn func() (interface{}, error)) *CacheService {
	return &CacheService{
		cache:     make(map[string]interface{}),
		expiry:    expiry,
		refreshFn: refreshFn,
	}
}

// GetValue возвращает значение из кэша или обновляет его при необходимости
func (cs *CacheService) GetValue(key string) (interface{}, error) {
	cs.mutex.RLock()
	value, found := cs.cache[key]
	cs.mutex.RUnlock()

	if found {
		return value, nil
	}

	// Если значение отсутствует или просрочено, обновляем кэш
	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	// Перепроверяем значение после получения блокировки
	value, found = cs.cache[key]
	if !found {
		// Обновляем кэш, вызывая функцию обновления
		newValue, err := cs.refreshFn()
		if err != nil {
			return nil, err
		}

		// Обновляем кэш
		cs.cache[key] = newValue

		// Устанавливаем таймер сброса значения через expiry
		time.AfterFunc(cs.expiry, func() {
			cs.mutex.Lock()
			defer cs.mutex.Unlock()
			delete(cs.cache, key)
		})

		return newValue, nil
	}

	return value, nil
}

// HTMLTemplate представляет HTML-шаблон
var HTMLTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Data View</title>
</head>
<body>
	<h1>Data View for ID {{.ID}}</h1>
	<p>Data: {{.Data}}</p>
</body>
</html>
`

// ViewData представляет данные для отображения на странице
type ViewData struct {
	ID   string
	Data interface{}
}

func main() {
	cache := NewCacheService(5*time.Minute, refreshFunction)

	// Создаем шаблон из строки
	tmpl, err := template.New("view").Parse(HTMLTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
			return
		}

		value, err := cache.GetValue(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}

		// Создаем данные для отображения на странице
		viewData := ViewData{
			ID:   id,
			Data: value,
		}

		// Отображаем шаблон с данными
		err = tmpl.Execute(w, viewData)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	})

	http.ListenAndServe(":8081", nil)
}
