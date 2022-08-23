package app

import (
	"context"
	"github.com/gorilla/mux"
	"graceful-shutdown/internal/domain/something/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CreateHttpServer(
	ctx context.Context,
	address string,
	handler handler.SomethingHandler,
) {
	router := buildRouter(handler)

	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println("http server shutdown")
		}
	}()
	log.Println("server listening", address)

	sig := <-termChan
	log.Println(sig, "received, shutdown process initiated")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Println("http server shutdown error:", err)
	}

	log.Println("server exited properly")
}

func buildRouter(handler handler.SomethingHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Handle).Methods("GET")

	return router
}
