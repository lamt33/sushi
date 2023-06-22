package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/lamt3/sushi/tuna/common/logger"
	"github.com/lamt3/sushi/tuna/common/web"
)

var (
	healthy int32
)

func Start() {
	pg := initPG()
	defer pg.Close()

	cache := initCache()
	defer cache.Close()

	wrappedH := initializeAppComponents(AppTemplate{
		PG:    pg,
		Cache: cache,
	})

	srv := web.Server(wrappedH)
	startGracefulShutdownServer(srv)
}

func startGracefulShutdownServer(srv *http.Server) {
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Info("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("Could not gracefully shutdown the server: %v\n", err)
			os.Exit(1)
		}
		close(done)
	}()
	logger.Info("Listening on %s", srv.Addr)
	atomic.StoreInt32(&healthy, 1)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Could not listen on %s: %v\n", srv.Addr, err)
		os.Exit(1)
	}

	<-done
	logger.Info("Server stopped")
}
