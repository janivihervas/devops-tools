package main

import (
	"net/http"

	"os"

	"fmt"

	"github.com/janivihervas/devops-tools/bar"
	"github.com/janivihervas/devops-tools/log"
	"github.com/pkg/errors"
)

func main() {
	l := log.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		l.LogInfo(r.Method, r.URL.Path, r.URL.RawQuery)
		_, err := w.Write([]byte("Bar version " + bar.Version))
		if err != nil {
			http.Error(w, fmt.Sprintf("%+v", errors.Wrap(err, "bar")), http.StatusInternalServerError)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	l.LogInfo("Starting bar server on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		l.LogError(fmt.Sprintf("%+v", errors.Wrap(err, "bar")))
	}
}
