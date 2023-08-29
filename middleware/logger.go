package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		// Sebelum handler dijalankan
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next(w, r, ps)

		// Setelah handler selesai dijalankan
		log.Printf(
			"Completed %s %s in %v",
			r.Method, r.URL.Path, time.Since(start),
		)
	}
}
