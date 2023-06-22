package server

import (
	"log"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/lamt3/sushi/tuna/common/cache"
	"github.com/lamt3/sushi/tuna/common/db"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/account"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/handlers"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/repo"
	"github.com/rs/cors"
)

type AppTemplate struct {
	PG    *db.OPPG
	Cache cache.ICache
}

func initializeAppComponents(app AppTemplate) http.HandlerFunc {

	acctRepo := repo.NewPGAccountRepo(app.PG)
	acctSvc := account.NewAccountSvc(acctRepo, app.Cache)
	accountHandler := handlers.NewAccountHandler(&acctSvc)

	return createWrapper(createRoutes(accountHandler))

}

func createRoutes(ah *handlers.AccountHandler) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/v1/user/account", ah.LoginUser)

	return router
}

func createWrapper(router *httprouter.Router) http.HandlerFunc {

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

func initPG() *db.OPPG {
	return db.OPPGBuilder().
		PrimaryConn("ABC", func(d *sqlx.DB) {
			d.SetMaxOpenConns(15)
			d.SetMaxIdleConns(15)
			d.SetConnMaxLifetime(2 * time.Minute)
		}).SecondaryConn([]string{"abc", "bcd"},
		func(d *sqlx.DB) {
			d.SetMaxOpenConns(15)
			d.SetMaxIdleConns(15)
			d.SetConnMaxLifetime(2 * time.Minute)
		}).
		Build()

}

func initCache() cache.ICache {
	return cache.NewRedisCache(
		&redis.Options{
			Addr:         "8080",
			Password:     "password",
			DB:           1,
			MaxRetries:   3,
			PoolSize:     10,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		},
	)
}
