package web

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/julienschmidt/httprouter"
	"github.com/lamt3/sushi/tuna/common/logger"
)

type Middleware func(h httprouter.Handle) httprouter.Handle

type MW struct {
	DD          *statsd.Client
	MiddleWares []Middleware
}

func (mw *MW) Add(h httprouter.Handle) httprouter.Handle {
	for _, m := range mw.MiddleWares {
		h = m(h)
	}
	return h
}

func (mw *MW) UseDD() {
	a := func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			startTime := time.Now()
			logger.Info("in dd mw")
			h(w, r, ps)
			// Calculate the duration
			duration := time.Since(startTime)

			logger.Info("in dd mw2")

			// Send the metric to Datadog
			err := mw.DD.Timing("api.request.duration", duration, []string{"route:" + r.URL.Path}, 1)
			if err != nil {
				logger.Error("error in dd mw %s", err)
			}
			err = mw.DD.Incr("api.request.count", []string{"route:" + r.URL.Path}, 1)
			if err != nil {
				logger.Error("error in dd mw %s", err)
			}
		}
	}
	mw.MiddleWares = append(mw.MiddleWares, a)
}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.SetOutput(os.Stdout) // logs go to Stderr by default
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r) // call ServeHTTP on the original handler

	})
}

func DatadogAPIMiddleware(h httprouter.Handle, statsdClient *statsd.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		startTime := time.Now()
		h(w, r, ps)
		// Calculate the duration
		duration := time.Since(startTime)

		// Send the metric to Datadog
		statsdClient.Timing("api.request.duration", duration, []string{"route:" + r.URL.Path}, 1)
	}
}
