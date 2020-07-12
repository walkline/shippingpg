package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/clientapi"
)

type ReceivingHandler struct {
	repo   clientapi.KeyValueRepository
	logger log.Logger
}

func NewReceivingHandler(repo clientapi.KeyValueRepository, l log.Logger) *ReceivingHandler {
	return &ReceivingHandler{
		repo:   repo,
		logger: l,
	}
}

func (h ReceivingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(t time.Time) {
		h.logger.Log(
			"path", r.URL.Path,
			"took", time.Since(t),
		)
	}(time.Now())

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.repo.FindByKey(context.Background(), keys[0])
	if err != nil {
		h.logger.Log("can't load port", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, ok := p.Value.(json.RawMessage)
	if !ok {
		h.logger.Log("can't cast value to json.RawMessage")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
