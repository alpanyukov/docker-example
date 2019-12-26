package main

import (
	"my/webapp/handlers"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	hh := handlers.NewHandlers(mux)

	hh.AddHandler("/hello", handlers.NewHelloHandler())
	hh.AddHandler("/counter", handlers.NewCounterMessaging())

	rh, exists := os.LookupEnv("REDIS_URL")
	if exists {
		redisMessaging := handlers.NewDBCounterMessaging(rh)
		defer redisMessaging.Conn.Close()
		hh.AddHandler("/counter-redis", redisMessaging)
	}

	fileCounterMessaging := handlers.NewFileCounterMessaging()
	defer fileCounterMessaging.File.Close()
	hh.AddHandler("/counter-file", fileCounterMessaging)

	linksMessaging := handlers.NewLinksMessaging(hh.Links...)
	hh.AddHandler("/", linksMessaging)

	http.ListenAndServe(":3000", mux)
}
