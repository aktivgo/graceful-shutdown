package impl

import (
	"graceful-shutdown/internal/domain/something/service"
	"net/http"
	"sync"
)

type handler struct {
	wg *sync.WaitGroup

	service service.SomethingService
}

func NewHandler(
	wg *sync.WaitGroup,
	service service.SomethingService,
) *handler {
	return &handler{
		wg:      wg,
		service: service,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	h.wg.Add(3)

	go func() {
		if err := h.service.Job1(h.wg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	go func() {
		if err := h.service.Job2(h.wg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	go func() {
		if err := h.service.Job3(h.wg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
}
