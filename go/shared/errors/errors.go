package errors

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func notFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}
