package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Handler interface {
	InitRoutes(e *gin.Engine)
}

type Server struct {
	httpSrv *http.Server
	handler Handler
	Engine  *gin.Engine
}

func NewServer(handler Handler) *Server {
	return &Server{
		handler: handler,
		Engine:  gin.Default(),
	}
}

func (s *Server) Serve(g *errgroup.Group, ctx context.Context, port string) {
	s.handler.InitRoutes(s.Engine)

	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.Engine,
	}

	g.Go(func() error {
		<-ctx.Done()
		s.shutdown()

		return nil
	})

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Could not start listener ", err)
	}
}

func (s *Server) shutdown() {
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
	}
	cancel()
}
