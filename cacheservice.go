package main

import (
	"fmt"
	"sync"
	"time"
)

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

// Пример использования:

func main() {
	cache := NewCacheService(5*time.Minute, refreshFunction)
	value, err := cache.GetValue("example_key")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Value:", value)
}

// refreshFunction обновляет значение, например, считывая из базы данных
func refreshFunction() (interface{}, error) {
	// Здесь может быть логика обновления значения, например, запрос к БД
	return "example_value", nil
}
