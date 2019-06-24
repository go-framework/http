package server

import (
	"context"
	"net/http"
)

var (
	defaultServer *Server = NewServer(GetDefaultConfig())
)

func SetServer(srv *Server) {
	defaultServer = srv
}

func GetServer() *Server {
	return defaultServer
}

func GetConfig() *Config {
	return defaultServer.GetConfig()

}

func SetConfig(config *Config) error {
	return defaultServer.SetConfig(config)
}

func GetAddr() string {
	return defaultServer.GetAddr()
}

func GetHandler() http.Handler {
	return defaultServer.GetHandler()
}

func SetHandler(handler http.Handler) {
	defaultServer.SetHandler(handler)
}

func ListenAndServe() error {
	return defaultServer.ListenAndServe()
}

func ListenAndServeTLS(certFile, keyFile string) error {
	return defaultServer.ListenAndServeTLS(certFile, keyFile)
}

func GracefulShutdown(ctx context.Context) error {
	return defaultServer.GracefulShutdown(ctx)
}
