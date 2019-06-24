package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	*http.Server
	Config *Config `json:"http" yaml:"http"`
}

func NewServer(config *Config) *Server {

	srv := &Server{
		Config: config,
		Server: new(http.Server),
	}

	_ = srv.Update()

	return srv
}

func (srv Server) GetConfig() *Config {
	return srv.Config
}

func (srv *Server) SetConfig(config *Config) error {
	srv.Config = config

	return srv.Update()
}

func (srv *Server) Update() error {
	srv.Server.Addr = srv.Config.Addr

	return nil
}

func (srv Server) GetAddr() string {
	return srv.Server.Addr
}

func (srv Server) GetHandler() http.Handler {
	return srv.Server.Handler
}

func (srv *Server) SetHandler(handler http.Handler) {
	if handler == nil {
		return
	}
	srv.Server.Handler = handler
}

func (srv *Server) ListenAndServe() error {
	err := srv.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
	err := srv.Server.ListenAndServeTLS(certFile, keyFile)
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (srv Server) GracefulShutdown(ctx context.Context) error {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Waiting ...
	<-quit

	// Shutdown gracefully
	if err := defaultServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
