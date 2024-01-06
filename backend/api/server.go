package api

import (
	"net/http"

	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/ryanlaycock/minespeeder/domain/games"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type HTTPServer struct {
	s *http.Server
}

// Entry point to server.go
func (m *MineSpeederServer) newHTTPServer() *HTTPServer {
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	r.Use(middleware.OapiRequestValidator(swagger))

	HandlerFromMux(m, r)

	return &HTTPServer{
		s: &http.Server{
			Handler: r,
			Addr:    ":8080",
		},
	}
}

func (m *MineSpeederServer) ListenAndServe() {
	err := m.httpServer.s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type MineSpeederServer struct {
	gamesManager games.GamesManager
	httpServer  *HTTPServer
}

// Make sure we conform to ServerInterface

var _ ServerInterface = (*MineSpeederServer)(nil)

func NewMineSpeederServer(gamesManager games.GamesManager) *MineSpeederServer {
	m := &MineSpeederServer{
		gamesManager: gamesManager,
	}

	m.httpServer = m.newHTTPServer()

	return m
}
