package app

import (
	"context"
	"log"
	"sync"
)

const (
	Address = ":8080"
)

func Run() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	handler := BuildHandler(wg)

	CreateHttpServer(ctx, Address, handler)

	log.Println("waiting for running jobs to finish")
	wg.Wait()
	log.Println("jobs finished, exiting")
}
