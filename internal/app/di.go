package app

import (
	"graceful-shutdown/internal/domain/something/handler"
	somethingHandlerImpl "graceful-shutdown/internal/domain/something/handler/impl"
	somethingServiceImpl "graceful-shutdown/internal/domain/something/service/impl"
	"sync"
)

func BuildHandler(wg *sync.WaitGroup) handler.SomethingHandler {
	service := somethingServiceImpl.NewService()
	handler := somethingHandlerImpl.NewHandler(wg, service)

	return handler
}
