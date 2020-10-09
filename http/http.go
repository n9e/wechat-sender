package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/n9e/wechat-sender/config"
	"github.com/n9e/wechat-sender/http/middleware"
	"github.com/n9e/wechat-sender/http/render"
	"github.com/n9e/wechat-sender/http/router"
)

var srv = &http.Server{
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func Start() {
	render.Init()

	r := mux.NewRouter().StrictSlash(false)
	router.ConfigRoutes(r)

	n := negroni.New()
	n.Use(middleware.NewRecovery())

	n.UseHandler(r)

	srv.Addr = config.Get().HTTP.Listen
	srv.Handler = n

	go func() {
		fmt.Println("http.listening:", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listening %s occur error: %s\n", srv.Addr, err)
			os.Exit(3)
		}
	}()
}

// Shutdown http server
func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("cannot shutdown http server:", err)
		os.Exit(2)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		fmt.Println("shutdown http server timeout of 5 seconds.")
	default:
		fmt.Println("http server stopped")
	}
}
