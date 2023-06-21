package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/lamt3/sushi/tuna/common/db"
	"github.com/lamt3/sushi/tuna/common/logger"
	"github.com/lamt3/sushi/tuna/common/web"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/account"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/handlers"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/repo"
	"github.com/rs/cors"
)

var (
	healthy int32
)

func Start() {
	wrappedH := initialize()
	srv := web.Server(wrappedH)

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

func initialize() http.HandlerFunc {
	pgDB := db.PGBuilder().
		ConnectionString("abc").
		AddSettings(func(d *sqlx.DB) {
			d.SetMaxOpenConns(15)
			d.SetMaxIdleConns(15)
			d.SetConnMaxLifetime(2 * time.Minute)
		}).
		Build()

	defer pgDB.Close()

	acctRepo := repo.NewPGAccountRepo(pgDB)
	acctSvc := account.NewAccountSvc(acctRepo)
	accountHandler := handlers.NewAccountHandler(&acctSvc)

	return createRoutes(accountHandler)

}

func createRoutes(ah *handlers.AccountHandler) http.HandlerFunc {
	router := httprouter.New()
	router.POST("/api/v1/user/account", ah.LoginUser)

	wrappedCors := cors.
		New(cors.Options{
			AllowedOrigins: []string{
				"http://localhost:*",
			},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
			AllowedMethods:   []string{"HEAD", "GET", "POST", "PATCH", "DELETE"},
			//Debug: true, // Enable Debugging for testing, consider disabling in production
		}).
		Handler(router)

	wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(wrappedCors, w, r)
		log.Printf(
			"%s %s (code=%d dt=%s written=%d)", r.Method, r.URL, m.Code, m.Duration, m.Written,
		)
	})

	return wrappedH
}
