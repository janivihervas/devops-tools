package main

import (
	"net/http"

	"os"

	"fmt"

	"github.com/janivihervas/devops-tools/foo"
	"github.com/janivihervas/devops-tools/log"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

func main() {
	l := log.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		l.LogInfo(r.Method, r.URL.Path, r.URL.RawQuery)
		_, err := w.Write([]byte(fmt.Sprintf("Foo version %s, random stuff = %s", foo.Version, uuid.NewV4().String())))
		if err != nil {
			http.Error(w, fmt.Sprintf("%+v", errors.Wrap(err, "foo")), http.StatusInternalServerError)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	l.LogInfo("Starting foo server on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		l.LogError(fmt.Sprintf("%+v", errors.Wrap(err, "foo")))
	}
}
