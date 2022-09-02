package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Handler interface {
	InitRoutes(e *gin.Engine)
}

type Logger interface {
	Info(info string)
	Error(err error)
}

type Server struct {
	httpSrv *http.Server
	handler Handler
	Engine  *gin.Engine
	logger  Logger
}

func NewServer(handler Handler, logger Logger) *Server {
	return &Server{
		handler: handler,
		Engine:  gin.Default(),
		logger:  logger,
	}
}

func (s *Server) Serve(ctx context.Context, g *errgroup.Group, port string) {
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
		s.logger.Error(err)
	}
}

func (s *Server) shutdown() {
	s.logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Error(err)
	}
	cancel()
}
