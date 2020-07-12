package server

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/walkline/shippingpg/clientapi"
	"github.com/walkline/shippingpg/clientapi/kvimporting"
)

type HTTPServer struct {
	http.Server
}

func NewHTTPServer(
	addr string,
	logger log.Logger,
	portsImporter kvimporting.ReaderImporter,
	portsRepo clientapi.KeyValueRepository,
) *HTTPServer {
	mux := http.NewServeMux()
	mux.Handle("/v1/ports/import", NewImportingHandler(portsImporter, logger))
	mux.Handle("/v1/ports", NewReceivingHandler(portsRepo, logger))

	return &HTTPServer{
		Server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}
