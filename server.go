package api

import (
	"net/http"
	"os"

	"github.com/Guilherme-De-Marchi/ai-hub/api/handlers"
	"github.com/Guilherme-De-Marchi/ai-hub/api/middlewares"
	"golang.org/x/exp/slog"
)

type Server struct {
	addr   string
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewServer(addr string) *Server {
	return &Server{
		addr:   addr,
		mux:    http.NewServeMux(),
		logger: slog.New(slog.NewTextHandler(os.Stdout)),
	}
}

func (srv *Server) Start() (err error) {
	slog.Debug("started listening on address %v\n", srv.addr)
	defer slog.Debug("server stopped listening with error %v", err)

	srv.setRoutes()
	err = http.ListenAndServe(srv.addr, srv.mux)
	return err
}

func (srv *Server) setRoutes() {
	srv.registerRoute("/", http.FileServer(http.Dir("../www/public")))
	srv.registerRoute("/home", handlers.GetHome())
}

func (srv *Server) registerRoute(pattern string, handler http.Handler) {
	h := middlewares.LogRequest(srv.logger, handler)
	srv.mux.Handle(pattern, h)
}
