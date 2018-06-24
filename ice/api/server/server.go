package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Server implements HTTP server.
type Server struct {
	httpSrv *http.Server
	config  Config
	logger  *zap.Logger
}

// New creates a new Server instance. It will run HTTP server
// that serves the site configured by conf.
func New(conf Config) (*Server, error) {
	logger, err := newLogger()
	if err != nil {
		return nil, err
	}

	s := &Server{
		logger: logger,
		config: conf,
	}

	s.httpSrv = newHTTPServer(s, zap.NewStdLog(logger))

	return s, nil
}

// newLogger creates a logger
func newLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction(zap.AddCaller())
	return logger, err
}

// newHTTPServer creates a new http.Server instance configured by conf.
func newHTTPServer(s *Server, logger *log.Logger) *http.Server {
	return &http.Server{
		Addr:     s.config.Addr,
		ErrorLog: logger,
		Handler:  s.newHandler(),
	}
}

// Open opens the server
func (s *Server) Open() {
	go func() {
		s.logger.Info("http server started listening",
			zap.String("addr", s.httpSrv.Addr),
		)

		if err := s.httpSrv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Error("listen failed",
				zap.Error(err),
			)
		}
	}()
}

// Stop stops the server. This should be called before the
// process terminates, otherwise there would be some resources
// that are not released appropriately.
func (s *Server) Stop() {
	// shutdown http server gracefully
	s.logger.Info("trying to shutdown http server gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), s.config.GracefulTimeout*time.Second)
	defer cancel()
	err := s.httpSrv.Shutdown(ctx)

	if err != nil {
		s.logger.Error("failed to shutdown http server gracefully",
			zap.Error(err),
		)
	} else {
		s.logger.Info("succeeded to shutdown http server gracefully")
	}

	// flush logger
	s.logger.Info("flushing logger")
	s.logger.Sync()
}
