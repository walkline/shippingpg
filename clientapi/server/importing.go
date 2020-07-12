package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/clientapi/kvimporting"
)

type ImportingHandler struct {
	importer kvimporting.ReaderImporter
	logger   log.Logger
}

func NewImportingHandler(i kvimporting.ReaderImporter, l log.Logger) *ImportingHandler {
	return &ImportingHandler{
		importer: i,
		logger:   l,
	}
}

func (h ImportingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(t time.Time) {
		h.logger.Log(
			"path", r.URL.Path,
			"took", time.Since(t),
		)
	}(time.Now())

	if !(r.Method == http.MethodPost || r.Method == http.MethodPut) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	defer r.Body.Close()
	err := h.importer.Import(context.Background(), r.Body)
	if err != nil {
		h.logger.Log("can't import json to repo", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
