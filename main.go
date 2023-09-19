package main

import (
	"fmt"
	"net/http"

	"github.com/qmz/go-gin-blog/pkg/setting"
	"github.com/qmz/go-gin-blog/routers"
)

func main() {

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
