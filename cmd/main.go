package main

import (
	"fmt"
	"lesson49/internal/handlers"
	"net/http"
)

func main() {
	mux := &http.ServeMux{}
	handlers.NewWeatherHandler().Register(mux)

	srv := http.Server{
		Addr:    config.Addr(),
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
