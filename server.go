package main

import (
	"my/webapp/handlers"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	hh := handlers.NewHandlers(mux)

	// Обычный счетчик
	hh.AddHandler("/counter", handlers.NewCounterMessaging())

	// Redis
	rh, exists := os.LookupEnv("REDIS_URL")
	if exists {
		redisMessaging := handlers.NewDBCounterMessaging(rh)
		defer redisMessaging.Conn.Close()
		hh.AddHandler("/counter-redis", redisMessaging)
	}

	// Работа с файлом
	hh.AddHandler("/counter-file", handlers.NewFileCounterMessaging())

	// Просто ссылки
	linksMessaging := handlers.NewLinksMessaging(hh.Links...)
	hh.AddHandler("/", linksMessaging)

	http.ListenAndServe(":3000", mux)
}
