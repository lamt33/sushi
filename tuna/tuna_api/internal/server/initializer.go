package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/felixge/httpsnoop"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/lamt3/sushi/tuna/common/cache"
	"github.com/lamt3/sushi/tuna/common/db"
	"github.com/lamt3/sushi/tuna/common/logger"
	"github.com/lamt3/sushi/tuna/common/web"
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

	cli, err := statsd.New("dd-agent:8125")
	if err != nil {
		logger.Error("error %s", err)

	}
	//114a119a-755b-4377-ac7b-4da80b2e04b9

	mw := web.MW{
		DD:          cli,
		MiddleWares: []web.Middleware{},
	}

	mw.UseDD()

	// router.GET("/api/v1/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	web.WriteJSON(w, 200, "im alive!!")
	// })

	router.GET("/api/v1/health", mw.Add(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		web.WriteJSON(w, 200, "im alive!!")
	}))

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
		PrimaryConn(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "plum_pg", "5432", "postgres", "postgres", "postgres"),
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
			Addr:         "6379",
			Password:     "",
			DB:           0,
			MaxRetries:   3,
			PoolSize:     10,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	)
}
